package CONTROLLERS

import (
	"ROOMS/MIDDLEWARE"
	"github.com/gorilla/mux"
	"net/http"
)

func InitAllController(r *mux.Router) {

	//UsersController
	r.HandleFunc("/User/login", TokenHandler).Methods("POST")
	r.Handle("/user/getall", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(getalluser)))

	//others controller
	///
}
