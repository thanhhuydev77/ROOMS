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

	//BlocksController
	r.Handle("/block/get-block/{idowner}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(GetBlockByOwner))).Methods("GET")
	r.Handle("/block/create", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(CreateBlock))).Methods("POST")
	r.Handle("/block/update/{Idblock}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(UpdateBlock))).Methods("PUT")

	r.Handle("/block/delete/{id}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(DeleteBlock))).Methods("DELETE")
	r.Handle("/block/delete", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(DeleteBlocks))).Methods("DELETE")
	///
}
