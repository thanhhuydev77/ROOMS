package CONTROLLERS

import (
	"github.com/gorilla/mux"
)

func InitAllController(r *mux.Router) {

	//UsersController
	r.HandleFunc("/User/login", TokenHandler).Methods("POST")
	r.HandleFunc("/User/register", UserRegister).Methods("POST")
	//r.Handle("/user/getall", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(getalluser)))

	//others controller
	///
}
