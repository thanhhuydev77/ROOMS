package BUSINESS

import (
	"ROOMS/MODELS"
	"ROOMS/STATICS"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Getall() []MODELS.BILLS {
	bills := []MODELS.BILLS{}
	db, err := STATICS.Connectdatabase()
	// Query all users
	if db == nil {

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
	return bills
}

func GetBillById(id int) []MODELS.BILLS {
	bills := []MODELS.BILLS{}
	db, err := sql.Open("mysql", "root:tjmwjm824594@(104.197.241.11:3306)/ROOM_SCHEMA?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Query all users
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
	return bills
}

func CreateBill(bill MODELS.BILLS) int {
	db, err := sql.Open("mysql", "root:tjmwjm824594@(104.197.241.11:3306)/ROOM_SCHEMA?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Query all users
	query := "insert into BILLS(idRoom,dateCheckOut,totalPrice,isCheckedOut) VALUES(?,?,?,?)"
	rows, err := db.Query(query, bill.IdRoom, bill.DateCheckOut, bill.TotalPrice, bill.IsCheckedOut)
	if err != nil {
		return 0
	}
	panic(err)
	defer rows.Close()
	return 1
}
