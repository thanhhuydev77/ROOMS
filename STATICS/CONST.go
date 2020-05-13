package STATICS

import (
	"database/sql"
)

const (
	APP_KEY           = "thisissecreckeyyesitisreallyofcourcetrustmeitiskeyofthisapphahaha"
	CONNECTION_STRING = "root:tjmwjm824594@(104.197.241.11:3306)/ROOM_SCHEMA?parseTime=true"
	DRIVER_NAME       = "mysql"
	REDISURL          = "redis://localhost:6379"
)

func Connectdatabase() (*sql.DB, error) {
	db, err := sql.Open(DRIVER_NAME, CONNECTION_STRING)
	if err != nil {
		return nil, nil
	}
	if err := db.Ping(); err != nil {
		return nil, nil
	}
	return db, nil
}
