package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
)

func GetServiceById(db *sql.DB, Id int) ([]MODELS.GET_SERVICES_REQUEST, bool) {
	return DATABASE.GetServiceById(db, Id)
}
func DeleteService(db *sql.DB, id int) (bool, error) {
	return DATABASE.DeleteService(db, id)
}

func CreateService(db *sql.DB, services []MODELS.SERVICE_INPUT) (bool, error) {
	return DATABASE.CreateService(db, services)
}

func DeleteServices(db *sql.DB, servicesId []int) (bool, error) {
	return DATABASE.DeleteServices(db, servicesId)
}

func UpdateService(db *sql.DB, service MODELS.SERVICE_INPUT) (bool, error) {
	return DATABASE.UpdateService(db, service)
}
