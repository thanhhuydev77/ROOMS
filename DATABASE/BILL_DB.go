package DATABASE

import (
	"ROOMS/MODELS"
	"log"
)

func GetallBills() ([]MODELS.BILLS, error) {
	bills := []MODELS.BILLS{}
	db, err := connectdatabase()

	if db == nil {
		return nil, err
	}
	rows, err := db.Query(`SELECT * FROM BILLS`)
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
		bills = append(bills, u)
	}
	return bills, nil
}

func GetBillById(id int) ([]MODELS.BILLS, error) {
	bills := []MODELS.BILLS{}
	db, err := connectdatabase()
	if db == nil {
		return nil, err
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
		bills = append(bills, u)
	}
	return bills, nil
}

func CreateBill(bill MODELS.BILLS) (int, error) {
	db, err := connectdatabase()
	if db == nil {
		return 0, err
	}

	query := "insert into BILLS(idRoom,dateCheckOut,totalPrice,isCheckedOut) VALUES(?,?,?,?)"
	rows, err := db.Query(query, bill.IdRoom, bill.DateCheckOut, bill.TotalPrice, bill.IsCheckedOut)
	if err != nil {
		return 0, err
	}
	panic(err)
	defer rows.Close()
	return 1, nil
}
