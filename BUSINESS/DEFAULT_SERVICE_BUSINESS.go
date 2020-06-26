package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
)

func Get_all_default_service(db *sql.DB) ([]MODELS.GET_DEFAULT_SERVICES_REQUEST, bool) {
	return DATABASE.Get_all_default_service(db)
}
