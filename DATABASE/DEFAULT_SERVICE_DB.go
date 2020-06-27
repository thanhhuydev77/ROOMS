package DATABASE

import (
	"ROOMS/MODELS"
	"database/sql"
	"log"
)

func Get_all_default_service(db *sql.DB) ([]MODELS.GET_DEFAULT_SERVICES_REQUEST, bool) {
	var default_services []MODELS.GET_DEFAULT_SERVICES_REQUEST
	//db, err := connectdatabase()
	//// Query all users
	if db == nil {
		log.Print("can not connect to database!")
		return default_services, false
	}
	//defer db.Close()

	rows, err := db.Query("SELECT DS.*,U.name FROM DEFAULT_SERVICES as DS INNER JOIN UNITS as U on DS.idUnit = U.id")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var defaultServices MODELS.GET_DEFAULT_SERVICES_REQUEST
		err := rows.Scan(&defaultServices.Id, &defaultServices.NameService, &defaultServices.Price, &defaultServices.Description, &defaultServices.IdUnit, &defaultServices.UnitName)
		if err != nil {
			log.Fatal(err)
		}
		default_services = append(default_services, defaultServices)
	}
	defer rows.Close()
	return default_services, true
}
