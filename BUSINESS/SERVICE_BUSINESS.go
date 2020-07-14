package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
)

//get service by id
func GetServiceById(db *sql.DB, Id int) ([]MODELS.GET_SERVICES_REQUEST, bool) {
	return DATABASE.GetServiceById(db, Id)
}

//delete service with its id
func DeleteService(db *sql.DB, id int) (bool, error) {
	return DATABASE.DeleteService(db, id)
}

//create a new service
func CreateService(db *sql.DB, services []MODELS.SERVICE_INPUT) (bool, error) {
	return DATABASE.CreateService(db, services)
}

//delete many services
func DeleteServices(db *sql.DB, servicesId []int) (bool, error) {
	return DATABASE.DeleteServices(db, servicesId)
}

//update a service
func UpdateService(db *sql.DB, service MODELS.SERVICE_INPUT) (bool, error) {
	return DATABASE.UpdateService(db, service)
}
