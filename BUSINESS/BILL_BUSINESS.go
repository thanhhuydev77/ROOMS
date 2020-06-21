package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	_ "github.com/go-sql-driver/mysql"
)

//get bill information by id
func GetBillById(id int) (MODELS.BILLS, bool, error) {
	bill, ok, err := DATABASE.GetBillById(id)
	return bill, ok, err
}

//get bill detail by id
func GetBillDetailById(id int) ([]MODELS.BILL_DETAILS, bool, error) {
	billDetail, ok, err := DATABASE.GetBillDetailById(id)
	return billDetail, ok, err
}

//func CreateBill(bill MODELS.BILLS) int {
//	result, _ := DATABASE.CreateBill(bill)
//	return result
//}

//create a new bill and bill detail
func CreateBill(CCR MODELS.CREATE_UPDATE_BILL_REQUEST) (int, error) {
	return DATABASE.CreateBill(CCR)
}
