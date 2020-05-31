package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
)

func GetCustomers(userId int) ([]MODELS.CUSTOMER, bool, error) {
	return DATABASE.GetCustomers(userId)
}

func GetCustomersByUserId(userId int) ([]MODELS.CUSTOMER_GET, bool, error) {
	return DATABASE.GetCustomersByUserId(userId)
}

func CreateCustomer(c MODELS.CUSTOMER_INPUT) (bool, error) {
	return DATABASE.CreateCustomer(c)
}

func DeleteCustomer(idCustomer int) (bool, error) {
	return DATABASE.DeleteCustomer(idCustomer)
}

func DeleteManyCustomers(ids []int) (bool, error) {
	return DATABASE.DeleteManyCustomers(ids)
}

func UpdateCustomer(c MODELS.CUSTOMER_UPDATE) (bool, error) {
	return DATABASE.UpdateCustomer(c)
}

func SelectNameRoom(idUser int) ([]string, bool, error) {
	return DATABASE.SelectNameRoom(idUser)
}
