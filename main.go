package main

import (
	"ROOMS/CONTROLLERS"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		print(w, "getting started!")
	}).Methods("GET")

	CONTROLLERS.InitUserController(r)

	http.ListenAndServe(":8080", r)
}
