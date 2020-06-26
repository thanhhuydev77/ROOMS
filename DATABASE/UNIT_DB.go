package DATABASE

import (
	"ROOMS/MODELS"
	"database/sql"
	"log"
)

func GetAllUnits(db *sql.DB) ([]MODELS.UNIT, bool) {
	var units []MODELS.UNIT
	//db, err := connectdatabase()
	// Query all users

	if db == nil {

		log.Print("can not connect to database!")
		return units, false
	}
	defer db.Close()

	rows, err := db.Query("select * from UNITS")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var unit MODELS.UNIT
		err := rows.Scan(&unit.Id, &unit.Name, &unit.Description)
		if err != nil {
			log.Fatal(err)
		}
		units = append(units, unit)
	}
	defer rows.Close()
	return units, true
}
