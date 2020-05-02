package main

import (
	"ROOMS/CONTROLLERS"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	r := NewRouter()
	CONTROLLERS.InitAllController(r)
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
