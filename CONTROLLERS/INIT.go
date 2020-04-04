package CONTROLLERS

import (
	"ROOMS/MIDDLEWARE"
	"net/http"

	"github.com/gorilla/mux"
)

func InitAllController(r *mux.Router) {

	//UsersController
	r.HandleFunc("/user/login", UserLogin).Methods("POST")
	r.HandleFunc("/user/register", UserRegister).Methods("POST")
	r.HandleFunc("/user/get-all-username", GetallUserName).Methods("GET")
	r.Handle("/user/validate", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(MIDDLEWARE.ValidateToken))).Methods("GET")
	r.Handle("/user/get-user", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(GetUser))).Methods("GET")
	r.Handle("/user/get-user/{Id}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(GetUser))).Methods("GET")
	//RoomsController
	r.Handle("/block/getblockbyowner/{idowner}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(GetBlockByOwner))).Methods("GET")
	///
}
