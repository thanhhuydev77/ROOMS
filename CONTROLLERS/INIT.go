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
	r.HandleFunc("/User/getallusername", GetallUserName).Methods("GET")

	//RoomsController
	r.Handle("/Block/getblockbyowner/{idowner}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(GetBlockByOwner))).Methods("GET")
	///
}
