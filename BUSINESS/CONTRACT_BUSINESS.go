package BUSINESS

import (
	"ROOMS/MODELS"
	"ROOMS/STATICS"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//func Getall() []MODELS.BILLS {
//	bills := []MODELS.BILLS{}
//	db, err := STATICS.Connectdatabase()
//	// Query all users
//	if db == nil {
//
//	}
//	rows, err := db.Query(`SELECT * FROM BILLS`)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		u := MODELS.BILLS{}
//
//		err := rows.Scan(&u.Id, &u.IdRoom, &u.DateCheckOut, &u.TotalPrice, &u.IsCheckedOut)
//		if err != nil {
//			log.Fatal(err)
//		}
//		bills = append(bills, u)
//	}
//	return bills
//}
//
//func GetBillById(id int) []MODELS.BILLS {
//	bills := []MODELS.BILLS{}
//	db, err := sql.Open("mysql", "root:tjmwjm824594@(104.197.241.11:3306)/ROOM_SCHEMA?parseTime=true")
//	if err != nil {
//		log.Fatal(err)
//	}
//	if err := db.Ping(); err != nil {
//		log.Fatal(err)
//	}
//
//	// Query all users
//	query := "SELECT * FROM BILLS where id = ?"
//	rows, err := db.Query(query, id)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		u := MODELS.BILLS{}
//
//		err := rows.Scan(&u.Id, &u.IdRoom, &u.DateCheckOut, &u.TotalPrice, &u.IsCheckedOut)
//		if err != nil {
//			log.Fatal(err)
//		}
//		bills = append(bills, u)
//	}
//	return bills
//}

func CreateContract(CCR MODELS.CREATE_CONTRACT_REQUEST) bool {
	db, err := STATICS.Connectdatabase()

	if err != nil {
		log.Fatal("Can't connet to database")
		return false
	}
	defer db.Close()

	// Query all users

	ctx := context.Background()
	tx, err1 := db.BeginTx(ctx, nil)
	if err1 != nil {
		log.Fatal(err)
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO CONTRACTS (idRoom, idOwner, idSlave, startDate, endDate, circlePay, deposit, dayPay, note, idBlock) VALUES (?,?,?,?,?,?,?,?,?,?)`, CCR.IdRoom, CCR.IdOwner, CCR.IdSlave, CCR.StartDate, CCR.EndDate, CCR.CirclePay, CCR.Deposit, CCR.DayPay, CCR.Note, CCR.IdBlock)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return false
	}

	// The next query is handled similarly
	for _, vals := range CCR.UserRooms {
		_, err = tx.ExecContext(ctx, `INSERT INTO USER_ROOM (idUser, idRoom) VALUES(?,?)`, vals.IdUser, vals.IdRoom)
		if err != nil {
			// Incase we find any error in the query execution, rollback the transaction
			tx.Rollback()
			return false
		}
	}

	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = nil
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	//defer rows.Close()
	return true
}
