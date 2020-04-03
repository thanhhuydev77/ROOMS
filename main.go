package main

import (
	"ROOMS/CONTROLLERS"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	CONTROLLERS.InitAllController(r)
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8001", handler)
}
