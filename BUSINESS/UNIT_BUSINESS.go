package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
)

func GetAllUnits(db *sql.DB) ([]MODELS.UNIT, bool) {
	return DATABASE.GetAllUnits(db)
}
