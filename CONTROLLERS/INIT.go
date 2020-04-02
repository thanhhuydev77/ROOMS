package CONTROLLERS

import (
	"ROOMS/MIDDLEWARE"
	"github.com/gorilla/mux"
	"net/http"
)

func InitAllController(r *mux.Router) {

	//UsersController
	r.HandleFunc("/User/login", TokenHandler).Methods("POST")
	r.HandleFunc("/User/register", UserRegister).Methods("POST")
	r.HandleFunc("/User/get-all-username", GetallUserName).Methods("GET")
	r.Handle("/User/get-user", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(GetUser))).Methods("GET")
	r.Handle("/User/get-user/{Id}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(GetUser))).Methods("GET")
	//RoomsController
	r.Handle("/Block/getblockbyowner/{idowner}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(GetBlockByOwner))).Methods("GET")
	///
}
