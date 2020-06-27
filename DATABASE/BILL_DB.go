package DATABASE

import (
	"ROOMS/MODELS"
	"context"
	"database/sql"
	"fmt"
	"log"
)

func GetBillByBlockId(db *sql.DB, idBlock int) ([]MODELS.BILLS, bool, error) {
	bills := []MODELS.BILLS{}
	//db, err := connectdatabase()
	//
	if db == nil {
		return nil, false, fmt.Errorf("can not connect db")
	}
	rows, err := db.Query(`select * from BILLS where idRoom in(select id from ROOMS where idBlock = ?)`, idBlock)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()
	for rows.Next() {
		u := MODELS.BILLS{}

		err := rows.Scan(&u.Id, &u.IdRoom, &u.DateCheckOut, &u.TotalPrice, &u.IsCheckedOut)
		if err != nil {
			log.Fatal(err)
		}
		bills = append(bills, u)
	}
	return bills, true, nil
}

func GetBillById(db *sql.DB, id int) (MODELS.BILLS, bool, error) {
	//db, err := connectdatabase()
	if db == nil {
		return MODELS.BILLS{}, false, fmt.Errorf("can not connect db")
	}

	// Query all bills with id = id
	query := "SELECT * FROM BILLS where id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		u := MODELS.BILLS{}
		err := rows.Scan(&u.Id, &u.IdRoom, &u.DateCheckOut, &u.TotalPrice, &u.IsCheckedOut)
		if err != nil {
			log.Fatal(err)
		}
		return u, true, nil
	}
	return MODELS.BILLS{}, false, nil
}

func GetBillDetailById(db *sql.DB, id int) ([]MODELS.BILL_DETAILS, bool, error) {
	listbilldt := []MODELS.BILL_DETAILS{}

	//db, err := connectdatabase()
	if db == nil {
		return nil, false, fmt.Errorf("can not connect db")
	}

	// Query all bills with id = id
	query := "SELECT * FROM BILL_DETAILS where idBill = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		u := MODELS.BILL_DETAILS{}
		err := rows.Scan(&u.Id, &u.IdBill, &u.IdService, &u.Amount, &u.TotalPrice)
		if err != nil {
			log.Fatal(err)
		}
		listbilldt = append(listbilldt, u)
	}
	return listbilldt, true, nil
}

func CreateBill(db *sql.DB, bill MODELS.CREATE_UPDATE_BILL_REQUEST) (int, error) {
	//db, err := connectdatabase()
	//if db == nil {
	//	return 0, err
	//}
	if db == nil {
		log.Print("can not connect to database!")
		return 0, fmt.Errorf("can not connect db")
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	query := "insert into BILLS(idRoom,dateCheckOut,totalPrice,isCheckedOut) VALUES(?,?,?,?)"
	_, errinsert := db.Query(query, bill.IdRoom, bill.DateCheckOut, bill.TotalPrice, bill.IsCheckedOut)
	if errinsert != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return 0, err
	}
	queryid := "select id from BILLS where idRoom = ? and dateCheckOut = ?"
	rowsid, errid := db.Query(queryid, bill.IdRoom, bill.DateCheckOut)
	defer rowsid.Close()
	if errid != nil {
		tx.Rollback()
		return 0, errid
	}
	id := 0
	for rowsid.Next() {
		errscan := rowsid.Scan(&id)
		if errscan != nil || id == 0 {
			tx.Rollback()
			return 0, errid
		}
	}
	for _, val := range bill.BillDetail {
		query := "INSERT INTO BILL_DETAILS(IdBill,IdService,Amount,TotalPrice) VALUES (?,?,?,?)"
		_, errinsert := db.Query(query, id, val.IdService, val.Amount, val.TotalPrice)
		if errinsert != nil {
			tx.Rollback()
			return 0, err
		}
	}
	errcmt := tx.Commit()
	if errcmt != nil {
		log.Printf("err while commit :", errcmt.Error())
		return 0, nil
	}
	return 1, nil
}

func UpdateBill(db *sql.DB, c MODELS.CREATE_UPDATE_BILL_REQUEST) (bool, error) {
	//db, err := connectdatabase()
	//if err != nil {
	//	log.Print("can not connect to database!")
	//	return false, err
	//}
	//defer db.Close()
	if db == nil {
		log.Print("can not connect to database!")
		return false, fmt.Errorf("can not connect db")
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("UPDATE BILLS SET dateCheckOut = ?, totalPrice = ?, isCheckedOut = ? WHERE id = ?", c.DateCheckOut, c.TotalPrice, c.IsCheckedOut, c.Id)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return false, err
	}

	_, err2 := db.Exec("DELETE FROM BILL_DETAILS WHERE idBill = ?", c.Id)
	if err2 != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return false, err
	}
	for _, vals := range c.BillDetail {
		_, err4 := tx.ExecContext(ctx, `INSERT INTO BILL_DETAILS(idBill, idService,amount,totalPrice) VALUES(?,?,?,?)`, c.Id, vals.IdService, vals.Amount, vals.TotalPrice)
		if err4 != nil {
			// Incase we find any error in the query execution, rollback the transaction
			tx.Rollback()
			return false, err
		}
	}
	err = nil
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return true, nil
}

func DeleteBill(db *sql.DB, idbill int) (bool, error) {
	//db, err := connectdatabase()
	//if err != nil {
	//	log.Fatalln(err)
	//	return false, err
	//}
	//defer db.Close()

	if db == nil {
		log.Print("can not connect to database!")
		return false, fmt.Errorf("can not connect db")
	}

	rs, err := db.Exec(`DELETE FROM BILLS WHERE id = ?`, idbill)

	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	num, err := rs.RowsAffected()
	if num == 0 {
		return false, err
	}
	return true, nil
}
