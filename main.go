package main

import (
	"ROOMS/CONTROLLERS"
	"ROOMS/DATABASE"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	r := NewRouter()
	Redis, _ := CONTROLLERS.NewStorage(DATABASE.REDISURL)
	app := &CONTROLLERS.ApiDB{
		Db: DATABASE.GetDbInstance(),
	}
	CONTROLLERS.InitAllController(*app, r, Redis)
	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":8001", handler)
}
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Server CSS, JS & Images Statically.
	router.
		PathPrefix("/public/").
		Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("."+"/public/"))))
	return router
}

func RecordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID); err != nil {
		return
	}
	return
}
