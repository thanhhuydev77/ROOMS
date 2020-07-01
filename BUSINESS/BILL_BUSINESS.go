package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//get bill information by id
func GetBillById(db *sql.DB, id int) (MODELS.BILLS, bool, error) {
	bill, ok, err := DATABASE.GetBillById(db, id)
	return bill, ok, err
}

//get bill detail by id
func GetBillDetailById(db *sql.DB, id int) ([]MODELS.BILL_DETAILS, bool, error) {
	billDetail, ok, err := DATABASE.GetBillDetailById(db, id)
	return billDetail, ok, err
}
func GetBillByIdblock(db *sql.DB, id int) ([]MODELS.BILLS, bool, error) {
	//get room
	return DATABASE.GetBillByBlockId(db, id)
}

//update a bill with its id
func UpdateBill(db *sql.DB, c MODELS.CREATE_UPDATE_BILL_REQUEST) (bool, error) {
	return DATABASE.UpdateBill(db, c)
}

//create a new bill and bill detail
func CreateBill(db *sql.DB, CCR MODELS.CREATE_UPDATE_BILL_REQUEST) (int, error) {
	return DATABASE.CreateBill(db, CCR)
}

func DeleteBill(db *sql.DB, idbill int) (bool, error) {
	return DATABASE.DeleteBill(db, idbill)
}
