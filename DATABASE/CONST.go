package DATABASE

import (
	"database/sql"
)

const (
	APP_KEY           = "thisissecreckeyyesitisreallyofcourcetrustmeitiskeyofthisapphahaha"
	CONNECTION_STRING = "sql12346166:B8rpnnyqRw@tcp(sql12.freemysqlhosting.net:3306)/sql12346166?parseTime=true"
	DRIVER_NAME       = "mysql"
	REDISURL          = "redis://localhost:6379"
)

func connectdatabase() (*sql.DB, error) {
	db, err := sql.Open(DRIVER_NAME, CONNECTION_STRING)
	if err != nil {
		return nil, nil
	}
	if err := db.Ping(); err != nil {
		return nil, nil
	}
	return db, nil
}
