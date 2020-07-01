package CONTROLLERS

import (
	"database/sql"
)

type ApiDB struct {
	Db *sql.DB
}
