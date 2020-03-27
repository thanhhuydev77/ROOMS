package main

import (
	"ROOMS/CONTROLLERS"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	CONTROLLERS.InitAllController(r)
	http.ListenAndServe(":8080", r)
}
