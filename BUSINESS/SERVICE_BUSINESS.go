package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
)

func GetServiceById(Id int) ([]MODELS.GET_SERVICES_REQUEST, bool) {
	return DATABASE.GetServiceById(Id)
}
func DeleteService(id int) (bool, error) {
	return DATABASE.DeleteService(id)
}

func CreateService(services []MODELS.SERVICE_INPUT) (bool, error) {
	return DATABASE.CreateService(services)
}

func DeleteServices(servicesId []int) (bool, error) {
	return DATABASE.DeleteServices(servicesId)
}

func UpdateService(service MODELS.SERVICE_INPUT) (bool, error) {
	return DATABASE.UpdateService(service)
}
