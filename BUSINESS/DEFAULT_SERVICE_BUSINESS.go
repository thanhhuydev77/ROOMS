package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
)

func Get_all_default_service() ([]MODELS.GET_DEFAULT_SERVICES_REQUEST, bool) {
	return DATABASE.Get_all_default_service()
}
