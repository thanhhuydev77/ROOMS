package CONTROLLERS

import (
	"database/sql"
)

//a interface to store a DB pointer
type ApiDB struct {
	Db *sql.DB
}
