package CONTROLLERS

import (
	M "ROOMS/MIDDLEWARE"

	"net/http"

	"github.com/gorilla/mux"
)

func InitAllController(r *mux.Router, storage *M.Storage) {

	//UsersController
	r.HandleFunc("/user/login", UserLogin).Methods("POST")
	r.HandleFunc("/user/register", UserRegister).Methods("POST")
	r.Handle("/user/get-all-username", M.Cached(storage, "10s", GetallUserName)).Methods("GET")
	r.Handle("/user/validate", M.AuthMiddleware(http.HandlerFunc(M.ValidateToken))).Methods("POST")
	r.Handle("/user/get-user", M.AuthMiddleware(M.Cached(storage, "10s", GetUser))).Methods("GET")
	r.Handle("/user/get-user/{Id}", M.AuthMiddleware(M.Cached(storage, "10s", GetUser))).Methods("GET")

	//BlocksController
	r.Handle("/block/get-block/{id}", M.AuthMiddleware(M.Cached(storage, "10s", GetBlockById))).Methods("GET")
	r.Handle("/block/get-block", M.AuthMiddleware(M.Cached(storage, "10s", GetBlock))).Methods("GET")
	r.Handle("/block/create", M.AuthMiddleware(http.HandlerFunc(CreateBlock))).Methods("POST")
	r.Handle("/block/update/{id}", M.AuthMiddleware(http.HandlerFunc(UpdateBlock))).Methods("PUT")

	r.Handle("/block/delete/{id}", M.AuthMiddleware(http.HandlerFunc(DeleteBlock))).Methods("DELETE")
	r.Handle("/block/delete-all", M.AuthMiddleware(http.HandlerFunc(DeleteBlocks))).Methods("POST")

	//RoomController
	r.Handle("/room/get-rooms", M.AuthMiddleware(M.Cached(storage, "10s", GetRoom))).Methods("GET")
	r.Handle("/room/create", M.AuthMiddleware(http.HandlerFunc(CreateRoom))).Methods("POST")
	r.Handle("/room/delete/{id}", M.AuthMiddleware(http.HandlerFunc(DeleteRoom))).Methods("DELETE")
	r.Handle("/room/delete-all", M.AuthMiddleware(http.HandlerFunc(DeleteRooms))).Methods("POST")
	r.Handle("/room/update/{id}", M.AuthMiddleware(http.HandlerFunc(UpdateRoom))).Methods("PUT")

	//UnitController
	r.Handle("/unit/get-units", M.AuthMiddleware(M.Cached(storage, "10s", GetAllUnit))).Methods("GET")

	//DefaultServiceController
	r.Handle("/default-service/get-default-services", M.AuthMiddleware(M.Cached(storage, "10s", Get_default_service))).Methods("GET")

	//ServiceController
	r.Handle("/service/get-services", M.AuthMiddleware(M.Cached(storage, "10s", GetService))).Methods("GET")
	r.Handle("/service/delete/{id}", M.AuthMiddleware(http.HandlerFunc(DeleteService))).Methods("DELETE")
	r.Handle("/service/create", M.AuthMiddleware(http.HandlerFunc(CreateService))).Methods("POST")
	r.Handle("/service/delete-all", M.AuthMiddleware(http.HandlerFunc(DeleteServices))).Methods("POST")
	r.Handle("/service/update/{id}", M.AuthMiddleware(http.HandlerFunc(UpdateService))).Methods("PUT")

	//uploadFile
	r.HandleFunc("/upload/userAvatar", UploadPicture).Methods("POST")

	//CustomerController
	r.Handle("/customer/get-customers", M.AuthMiddleware(http.HandlerFunc(GetCustomersByUserId))).Methods("GET")
	r.Handle("/customer/create", M.AuthMiddleware(http.HandlerFunc(CreateCustomer))).Methods("POST")
	r.Handle("/customer/delete/{id}", M.AuthMiddleware(http.HandlerFunc(DeleteCustomer))).Methods("DELETE")
	r.Handle("/customer/delete-all", M.AuthMiddleware(http.HandlerFunc(DeleteManyCustomers))).Methods("POST")
	r.Handle("/customer/update/{id}", M.AuthMiddleware(http.HandlerFunc(UpdateCustomer))).Methods("PUT")

	//ContractController
	r.Handle("/contract/get-contracts", M.AuthMiddleware(http.HandlerFunc(GetContract))).Methods("GET")
	r.Handle("/contract/create", M.AuthMiddleware(http.HandlerFunc(CreateContract))).Methods("POST")
	r.Handle("/contract/delete/{id}", M.AuthMiddleware(http.HandlerFunc(DeleteContract))).Methods("DELETE")
	r.Handle("/contract/delete-all", M.AuthMiddleware(http.HandlerFunc(DeleteAllContract))).Methods("POST")
	r.Handle("/contract/update/{id}", M.AuthMiddleware(http.HandlerFunc(UpdateContract))).Methods("PUT")
}
