package main

import (
	. "ROOMS/CONTROLLERS"
	"ROOMS/CONTROLLERS/cache"
	"ROOMS/DATABASE"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	r := NewRouter()
	Redis, _ := cache.NewStorage(DATABASE.REDISURL)
	app := &ApiDB{
		Db: DATABASE.GetDbInstance(),
	}
	defer app.Db.Close()
	fmt.Print("Server is running...")
	InitAllController(*app, r, Redis)
	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":8001", handler)

}
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Server CSS, JS & Images Statically.
	router.
		PathPrefix("/public/").
		Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("."+"/public/"))))
	return router
}
func InitAllController(a ApiDB, r *mux.Router, storage *cache.Storage) {

	//UsersController
	r.HandleFunc("/user/login", a.UserLogin).Methods("POST")
	r.HandleFunc("/user/register", a.UserRegister).Methods("POST")
	r.Handle("/user/get-all-username", cache.Cached(storage, "10s", a.GetallUserName)).Methods("GET")
	r.Handle("/user/validate", AuthMiddleware(http.HandlerFunc(ValidateToken))).Methods("POST")
	r.Handle("/user/get-user", AuthMiddleware(cache.Cached(storage, "10s", a.GetUser))).Methods("GET")
	r.Handle("/user/get-user/{Id}", AuthMiddleware(cache.Cached(storage, "10s", a.GetUser))).Methods("GET")

	//BlocksController
	r.Handle("/block/get-block/{id}", AuthMiddleware(cache.Cached(storage, "10s", a.GetBlockById))).Methods("GET")
	r.Handle("/block/get-block", AuthMiddleware(cache.Cached(storage, "10s", a.GetBlock))).Methods("GET")
	r.Handle("/block/create", AuthMiddleware(http.HandlerFunc(a.CreateBlock))).Methods("POST")
	r.Handle("/block/update/{id}", AuthMiddleware(http.HandlerFunc(a.UpdateBlock))).Methods("PUT")
	r.Handle("/block/delete/{id}", AuthMiddleware(http.HandlerFunc(a.DeleteBlock))).Methods("DELETE")
	r.Handle("/block/delete-all", AuthMiddleware(http.HandlerFunc(a.DeleteBlocks))).Methods("POST")

	//RoomController
	r.Handle("/room/get-rooms", AuthMiddleware(cache.Cached(storage, "10s", a.GetRoom))).Methods("GET")
	r.Handle("/room/get-rooms/{id}", AuthMiddleware(cache.Cached(storage, "10s", a.GetRoom))).Methods("GET")
	r.Handle("/room/get-rooms-dashboard", AuthMiddleware(cache.Cached(storage, "10s", a.GetRoomDB))).Methods("GET")
	r.Handle("/room/get-images", AuthMiddleware(cache.Cached(storage, "10s", a.GetRoomImage))).Methods("GET")
	r.Handle("/room/get-user-rent", AuthMiddleware(cache.Cached(storage, "10s", a.GetRoomUser))).Methods("GET")
	r.Handle("/room/create", AuthMiddleware(http.HandlerFunc(a.CreateRoom))).Methods("POST")
	r.Handle("/room/delete/{id}", AuthMiddleware(http.HandlerFunc(a.DeleteRoom))).Methods("DELETE")
	r.Handle("/room/delete-all", AuthMiddleware(http.HandlerFunc(a.DeleteRooms))).Methods("POST")
	r.Handle("/room/update/{id}", AuthMiddleware(http.HandlerFunc(a.UpdateRoom))).Methods("PUT")

	//UnitController
	r.Handle("/unit/get-units", AuthMiddleware(cache.Cached(storage, "10s", a.GetAllUnit))).Methods("GET")

	//DefaultServiceController
	r.Handle("/default-service/get-default-services", AuthMiddleware(cache.Cached(storage, "10s", a.Get_default_service))).Methods("GET")

	//ServiceController
	r.Handle("/service/get-services", AuthMiddleware(cache.Cached(storage, "10s", a.GetService))).Methods("GET")
	r.Handle("/service/delete/{id}", AuthMiddleware(http.HandlerFunc(a.DeleteService))).Methods("DELETE")
	r.Handle("/service/create", AuthMiddleware(http.HandlerFunc(a.CreateService))).Methods("POST")
	r.Handle("/service/delete-all", AuthMiddleware(http.HandlerFunc(a.DeleteServices))).Methods("POST")
	r.Handle("/service/update/{id}", AuthMiddleware(http.HandlerFunc(a.UpdateService))).Methods("PUT")

	//uploadFile
	r.HandleFunc("/upload/userAvatar", UploadPicture).Methods("POST")

	//CustomerController
	r.Handle("/customer/get-customers", AuthMiddleware(http.HandlerFunc(a.GetCustomersByUserId))).Methods("GET")
	r.Handle("/customer/create", AuthMiddleware(http.HandlerFunc(a.CreateCustomer))).Methods("POST")
	r.Handle("/customer/delete/{id}", AuthMiddleware(http.HandlerFunc(a.DeleteCustomer))).Methods("DELETE")
	r.Handle("/customer/delete-all", AuthMiddleware(http.HandlerFunc(a.DeleteManyCustomers))).Methods("POST")
	r.Handle("/customer/update/{id}", AuthMiddleware(http.HandlerFunc(a.UpdateCustomer))).Methods("PUT")

	//ContractController
	r.Handle("/contract/get-contracts", AuthMiddleware(http.HandlerFunc(a.GetContract))).Methods("GET")
	r.Handle("/contract/create", AuthMiddleware(http.HandlerFunc(a.CreateContract))).Methods("POST")
	r.Handle("/contract/delete/{id}", AuthMiddleware(http.HandlerFunc(a.DeleteContract))).Methods("DELETE")
	r.Handle("/contract/delete-all", AuthMiddleware(http.HandlerFunc(a.DeleteAllContract))).Methods("POST")
	r.Handle("/contract/update/{id}", AuthMiddleware(http.HandlerFunc(a.UpdateContract))).Methods("PUT")

	//BillControllers
	r.Handle("/bill/get-bills/{id}", AuthMiddleware(http.HandlerFunc(a.GetBillsbyblock))).Methods("GET")
	r.Handle("/bill/get-bill-by-id/{id}", AuthMiddleware(http.HandlerFunc(a.GetBills))).Methods("GET")
	r.Handle("/bill/create", AuthMiddleware(http.HandlerFunc(a.CreateBill))).Methods("POST")
	r.Handle("/bill/delete/{id}", AuthMiddleware(http.HandlerFunc(a.DeleteBill))).Methods("DELETE")
	//r.Handle("/contract/delete-all", AuthMiddleware(http.HandlerFunc(DeleteAllContract))).Methods("POST")
	r.Handle("/bill/update/{id}", AuthMiddleware(http.HandlerFunc(a.UpdateBill))).Methods("PUT")
}
