package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
)

//get a customer with its id
func GetCustomers(db *sql.DB, userId int) ([]MODELS.CUSTOMER, bool, error) {
	return DATABASE.GetCustomers(db, userId)
}

//get a customer with id user
func GetCustomersByUserId(db *sql.DB, userId int, page int, limit int) ([]MODELS.CUSTOMER_GET, bool, error, int) {
	return DATABASE.GetCustomersByUserId(db, userId, page, limit)
}

//create a new customer
func CreateCustomer(db *sql.DB, c MODELS.CUSTOMER_INPUT) (bool, error) {
	return DATABASE.CreateCustomer(db, c)
}

//delete a customer with id
func DeleteCustomer(db *sql.DB, idCustomer int) (bool, error) {
	return DATABASE.DeleteCustomer(db, idCustomer)
}

//delete many customers
func DeleteManyCustomers(db *sql.DB, ids []int) (bool, error) {
	return DATABASE.DeleteManyCustomers(db, ids)
}

//update a customer
func UpdateCustomer(db *sql.DB, c MODELS.CUSTOMER_UPDATE) (bool, error) {
	return DATABASE.UpdateCustomer(db, c)
}

//get name room
func SelectNameRoom(db *sql.DB, idUser int) ([]string, bool, error) {
	return DATABASE.SelectNameRoom(db, idUser)
}
