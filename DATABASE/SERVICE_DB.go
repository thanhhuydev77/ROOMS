package DATABASE

import (
	"ROOMS/MODELS"
	"database/sql"
	"log"
	"strings"
)

func GetServiceById(db *sql.DB, Id int) ([]MODELS.GET_SERVICES_REQUEST, bool) {
	var Services []MODELS.GET_SERVICES_REQUEST
	//db, err := connectdatabase()
	//// Query all users
	//if db == nil {
	//
	//	log.Print("can not connect to database!")
	//	return Services, false
	//}
	//defer db.Close()

	rows, err := db.Query("SELECT S.*, U.name nameUnit FROM SERVICES S INNER JOIN UNITS U ON S.idUnit = U.id WHERE idBlock = ?", Id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var services MODELS.GET_SERVICES_REQUEST
		err := rows.Scan(&services.Id, &services.NameService, &services.Price, &services.IdUnit, &services.Description, &services.IdBlock, &services.UnitName)
		if err != nil {
			log.Fatal(err)
		}
		Services = append(Services, services)
	}
	defer rows.Close()
	return Services, true
}

func DeleteService(db *sql.DB, id int) (bool, error) {
	//db, err := connectdatabase()
	//
	//if err != nil {
	//	log.Print("can not connect to database!")
	//	return false, err
	//}
	//defer db.Close()

	res, err := db.Exec(`delete from SERVICES where id = ?`, id)

	if err != nil {
		panic(err)
	}

	num, err := res.RowsAffected()
	m := int64(num)
	if m == 0 {
		return false, err
	}
	return true, nil
}

func CreateService(db *sql.DB, services []MODELS.SERVICE_INPUT) (bool, error) {

	//db, err := connectdatabase()
	//if err != nil {
	//	log.Print("can not connect to database!")
	//	return false, err
	//}
	//defer db.Close()

	sqlStr := "insert into SERVICES(nameService, price, idUnit, description, idBlock) values "
	vals := []interface{}{}

	for _, row := range services {
		sqlStr += "(?, ?, ?, ?, ?),"
		vals = append(vals, row.NameService, row.Price, row.IdUnit, row.Description, row.IdBlock)
	}

	sqlStr = strings.TrimSuffix(sqlStr, ",")

	stmt, _ := db.Prepare(sqlStr)
	res, err := stmt.Exec(vals...)

	if err != nil || res == nil {
		return false, err
	}

	return true, nil
}

func DeleteServices(db *sql.DB, servicesId []int) (bool, error) {
	//db, err := connectdatabase()
	//if err != nil {
	//	log.Print("can not connect to database!")
	//	return false, err
	//}
	//defer db.Close()

	args := make([]interface{}, len(servicesId))
	for i, id := range servicesId {
		args[i] = id
	}

	stmt := `delete from SERVICES where id in (?` + strings.Repeat(",?", len(args)-1) + `)`
	rows, err := db.Exec(stmt, args...)

	num, err := rows.RowsAffected()
	m := int64(num)
	if m == 0 {
		return false, err
	}
	return true, nil

	return true, nil
}

func UpdateService(db *sql.DB, service MODELS.SERVICE_INPUT) (bool, error) {
	//db, err := connectdatabase()
	//if err != nil {
	//	log.Print("can not connect to database!")
	//	return false, err
	//}
	//defer db.Close()

	rows, err := db.Exec("UPDATE SERVICES SET price = ?, idUnit = ?, description = ? WHERE id = ?",
		service.Price, service.IdUnit, service.Description, service.Id)

	num, err := rows.RowsAffected()
	m := int64(num)
	if m == 0 {
		return false, err
	}
	return true, nil
}
