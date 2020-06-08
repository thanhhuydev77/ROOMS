package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	_ "github.com/go-sql-driver/mysql"
)

func Getall() []MODELS.BILLS {
	bills, _ := DATABASE.GetallBills()
	return bills
}

func GetBillById(id int) []MODELS.BILLS {
	bills, _ := DATABASE.GetBillById(id)
	return bills
}

func CreateBill(bill MODELS.BILLS) int {
	result, _ := DATABASE.CreateBill(bill)
	return result
}
