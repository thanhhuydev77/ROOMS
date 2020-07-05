package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
)

//get all units
func GetAllUnits(db *sql.DB) ([]MODELS.UNIT, bool) {
	return DATABASE.GetAllUnits(db)
}
