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
	r.Handle("/user/validate", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(MIDDLEWARE.ValidateToken))).Methods("POST")
	r.Handle("/user/get-user", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(GetUser))).Methods("GET")
	r.Handle("/user/get-user/{Id}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(GetUser))).Methods("GET")

	//BlocksController
	r.Handle("/block/get-block/{id}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(GetBlockBYId))).Methods("GET")
	r.Handle("/block/get-block", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(GetBlock))).Methods("GET")
	r.Handle("/block/create", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(CreateBlock))).Methods("POST")
	r.Handle("/block/update/{id}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(UpdateBlock))).Methods("PUT")

	r.Handle("/block/delete/{id}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(DeleteBlock))).Methods("DELETE")
	r.Handle("/block/delete-all", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(DeleteBlocks))).Methods("POST")

	//RoomController
	r.Handle("/room/get-rooms/{idBlock}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(GetRoom))).Methods("GET")
	r.Handle("/room/create", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(CreateRoom))).Methods("POST")
	r.Handle("/room/delete/{id}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(DeleteRoom))).Methods("DELETE")
	r.Handle("/room/delete-all", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(DeleteRooms))).Methods("POST")
	r.Handle("/room/update/{id}", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(UpdateRoom))).Methods("PUT")
}
