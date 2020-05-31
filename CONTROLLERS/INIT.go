package CONTROLLERS

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitAllController(r *mux.Router, storage *Storage) {

	//UsersController
	r.HandleFunc("/user/login", UserLogin).Methods("POST")
	r.HandleFunc("/user/register", UserRegister).Methods("POST")
	r.Handle("/user/get-all-username", Cached(storage, "10s", GetallUserName)).Methods("GET")
	r.Handle("/user/validate", AuthMiddleware(http.HandlerFunc(ValidateToken))).Methods("POST")
	r.Handle("/user/get-user", AuthMiddleware(Cached(storage, "10s", GetUser))).Methods("GET")
	r.Handle("/user/get-user/{Id}", AuthMiddleware(Cached(storage, "10s", GetUser))).Methods("GET")

	//BlocksController
	r.Handle("/block/get-block/{id}", AuthMiddleware(Cached(storage, "10s", GetBlockById))).Methods("GET")
	r.Handle("/block/get-block", AuthMiddleware(Cached(storage, "10s", GetBlock))).Methods("GET")
	r.Handle("/block/create", AuthMiddleware(http.HandlerFunc(CreateBlock))).Methods("POST")
	r.Handle("/block/update/{id}", AuthMiddleware(http.HandlerFunc(UpdateBlock))).Methods("PUT")

	r.Handle("/block/delete/{id}", AuthMiddleware(http.HandlerFunc(DeleteBlock))).Methods("DELETE")
	r.Handle("/block/delete-all", AuthMiddleware(http.HandlerFunc(DeleteBlocks))).Methods("POST")

	//RoomController
	r.Handle("/room/get-rooms", AuthMiddleware(Cached(storage, "10s", GetRoom))).Methods("GET")
	r.Handle("/room/get-rooms/{id}", AuthMiddleware(Cached(storage, "10s", GetRoom))).Methods("GET")
	r.Handle("/room/get-rooms-dashboard", AuthMiddleware(Cached(storage, "10s", GetRoomDB))).Methods("GET")
	r.Handle("/room/get-images", AuthMiddleware(Cached(storage, "10s", GetRoomImage))).Methods("GET")
	r.Handle("/room/get-user-rent", AuthMiddleware(Cached(storage, "10s", GetRoomUser))).Methods("GET")
	r.Handle("/room/create", AuthMiddleware(http.HandlerFunc(CreateRoom))).Methods("POST")
	r.Handle("/room/delete/{id}", AuthMiddleware(http.HandlerFunc(DeleteRoom))).Methods("DELETE")
	r.Handle("/room/delete-all", AuthMiddleware(http.HandlerFunc(DeleteRooms))).Methods("POST")
	r.Handle("/room/update/{id}", AuthMiddleware(http.HandlerFunc(UpdateRoom))).Methods("PUT")

	//UnitController
	r.Handle("/unit/get-units", AuthMiddleware(Cached(storage, "10s", GetAllUnit))).Methods("GET")

	//DefaultServiceController
	r.Handle("/default-service/get-default-services", AuthMiddleware(Cached(storage, "10s", Get_default_service))).Methods("GET")

	//ServiceController
	r.Handle("/service/get-services", AuthMiddleware(Cached(storage, "10s", GetService))).Methods("GET")
	r.Handle("/service/delete/{id}", AuthMiddleware(http.HandlerFunc(DeleteService))).Methods("DELETE")
	r.Handle("/service/create", AuthMiddleware(http.HandlerFunc(CreateService))).Methods("POST")
	r.Handle("/service/delete-all", AuthMiddleware(http.HandlerFunc(DeleteServices))).Methods("POST")
	r.Handle("/service/update/{id}", AuthMiddleware(http.HandlerFunc(UpdateService))).Methods("PUT")

	//uploadFile
	r.HandleFunc("/upload/userAvatar", UploadPicture).Methods("POST")

	//CustomerController
	r.Handle("/customer/get-customers", AuthMiddleware(http.HandlerFunc(GetCustomersByUserId))).Methods("GET")
	r.Handle("/customer/create", AuthMiddleware(http.HandlerFunc(CreateCustomer))).Methods("POST")
	r.Handle("/customer/delete/{id}", AuthMiddleware(http.HandlerFunc(DeleteCustomer))).Methods("DELETE")
	r.Handle("/customer/delete-all", AuthMiddleware(http.HandlerFunc(DeleteManyCustomers))).Methods("POST")
	r.Handle("/customer/update/{id}", AuthMiddleware(http.HandlerFunc(UpdateCustomer))).Methods("PUT")

	//ContractController
	r.Handle("/contract/get-contracts", AuthMiddleware(http.HandlerFunc(GetContract))).Methods("GET")
	r.Handle("/contract/create", AuthMiddleware(http.HandlerFunc(CreateContract))).Methods("POST")
	r.Handle("/contract/delete/{id}", AuthMiddleware(http.HandlerFunc(DeleteContract))).Methods("DELETE")
	r.Handle("/contract/delete-all", AuthMiddleware(http.HandlerFunc(DeleteAllContract))).Methods("POST")
	r.Handle("/contract/update/{id}", AuthMiddleware(http.HandlerFunc(UpdateContract))).Methods("PUT")

}
