package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
)

func GetAllUnits() ([]MODELS.UNIT, bool) {
	return DATABASE.GetAllUnits()
}
