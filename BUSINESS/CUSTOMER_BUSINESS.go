package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
)

func GetCustomers(db *sql.DB, userId int) ([]MODELS.CUSTOMER, bool, error) {
	return DATABASE.GetCustomers(db, userId)
}

func GetCustomersByUserId(db *sql.DB, userId int) ([]MODELS.CUSTOMER_GET, bool, error) {
	return DATABASE.GetCustomersByUserId(db, userId)
}

func CreateCustomer(db *sql.DB, c MODELS.CUSTOMER_INPUT) (bool, error) {
	return DATABASE.CreateCustomer(db, c)
}

func DeleteCustomer(db *sql.DB, idCustomer int) (bool, error) {
	return DATABASE.DeleteCustomer(db, idCustomer)
}

func DeleteManyCustomers(db *sql.DB, ids []int) (bool, error) {
	return DATABASE.DeleteManyCustomers(db, ids)
}

func UpdateCustomer(db *sql.DB, c MODELS.CUSTOMER_UPDATE) (bool, error) {
	return DATABASE.UpdateCustomer(db, c)
}

func SelectNameRoom(db *sql.DB, idUser int) ([]string, bool, error) {
	return DATABASE.SelectNameRoom(db, idUser)
}
