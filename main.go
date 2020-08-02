package main

import (
	. "ROOMS/CONTROLLERS"
	c "ROOMS/CONTROLLERS/cache"
	"ROOMS/DATABASE"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := NewRouter()
	//init Redis server
	Redis, _ := c.NewStorage(DATABASE.REDISURL)
	app := &ApiDB{
		Db: DATABASE.GetDbInstance(),
	}
	defer app.Db.Close()
	fmt.Print("Server is running at port 8001...")
	InitAllController(*app, r, Redis)
	//allow all method CORS
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
func InitAllController(a ApiDB, r *mux.Router, storage *c.Storage) {

	//UsersController
	r.HandleFunc("/user/login", a.UserLogin).Methods("POST")
	r.HandleFunc("/user/register", a.UserRegister).Methods("POST")
	r.Handle("/user/get-all-username", c.Cached(storage, "10s", a.GetallUserName)).Methods("GET")
	r.Handle("/user/validate", AuthMW(http.HandlerFunc(ValidateToken))).Methods("POST")
	r.Handle("/user/get-user", AuthMW(http.HandlerFunc(a.GetUser))).Methods("GET")
	r.Handle("/user/get-user/{Id}", AuthMW(http.HandlerFunc(a.GetUser))).Methods("GET")

	//BlocksController
	r.Handle("/block/get-block/{id}", AuthMW(http.HandlerFunc(a.GetBlockById))).Methods("GET")
	r.Handle("/block/get-block", AuthMW(http.HandlerFunc(a.GetBlock))).Methods("GET")
	r.Handle("/block/create", AuthMW(http.HandlerFunc(a.CreateBlock))).Methods("POST")
	r.Handle("/block/update/{id}", AuthMW(http.HandlerFunc(a.UpdateBlock))).Methods("PUT")
	r.Handle("/block/delete/{id}", AuthMW(http.HandlerFunc(a.DeleteBlock))).Methods("DELETE")
	r.Handle("/block/delete-all", AuthMW(http.HandlerFunc(a.DeleteBlocks))).Methods("POST")

	//RoomController
	r.Handle("/room/get-rooms", AuthMW(http.HandlerFunc(a.GetRoom))).Methods("GET")
	r.Handle("/room/get-rooms/{id}", AuthMW(http.HandlerFunc(a.GetRoom))).Methods("GET")
	r.Handle("/room/get-rooms-dashboard", AuthMW(http.HandlerFunc(a.GetRoomDB))).Methods("GET")
	r.Handle("/room/get-images", AuthMW(http.HandlerFunc(a.GetRoomImage))).Methods("GET")
	r.Handle("/room/get-user-rent", AuthMW(http.HandlerFunc(a.GetRoomUser))).Methods("GET")
	r.Handle("/room/create", AuthMW(http.HandlerFunc(a.CreateRoom))).Methods("POST")
	r.Handle("/room/delete/{id}", AuthMW(http.HandlerFunc(a.DeleteRoom))).Methods("DELETE")
	r.Handle("/room/delete-all", AuthMW(http.HandlerFunc(a.DeleteRooms))).Methods("POST")
	r.Handle("/room/update/{id}", AuthMW(http.HandlerFunc(a.UpdateRoom))).Methods("PUT")

	//UnitController
	r.Handle("/unit/get-units", AuthMW(c.Cached(storage, "1h", a.GetAllUnit))).Methods("GET")

	//DefaultServiceController
	r.Handle("/default-service/get-default-services", AuthMW(c.Cached(storage, "1h", a.Get_default_service))).Methods("GET")

	//ServiceController
	r.Handle("/service/get-services", AuthMW(http.HandlerFunc(a.GetService))).Methods("GET")
	r.Handle("/service/delete/{id}", AuthMW(http.HandlerFunc(a.DeleteService))).Methods("DELETE")
	r.Handle("/service/create", AuthMW(http.HandlerFunc(a.CreateService))).Methods("POST")
	r.Handle("/service/delete-all", AuthMW(http.HandlerFunc(a.DeleteServices))).Methods("POST")
	r.Handle("/service/update/{id}", AuthMW(http.HandlerFunc(a.UpdateService))).Methods("PUT")

	//uploadFile
	r.HandleFunc("/upload/userAvatar", UploadPicture).Methods("POST")

	//CustomerController
	r.Handle("/customer/get-customers", AuthMW(http.HandlerFunc(a.GetCustomersByUserId))).Methods("GET")
	r.Handle("/customer/create", AuthMW(http.HandlerFunc(a.CreateCustomer))).Methods("POST")
	r.Handle("/customer/delete/{id}", AuthMW(http.HandlerFunc(a.DeleteCustomer))).Methods("DELETE")
	r.Handle("/customer/delete-all", AuthMW(http.HandlerFunc(a.DeleteManyCustomers))).Methods("POST")
	r.Handle("/customer/update/{id}", AuthMW(http.HandlerFunc(a.UpdateCustomer))).Methods("PUT")

	//ContractController
	r.Handle("/contract/get-contracts", AuthMW(http.HandlerFunc(a.GetContract))).Methods("GET")
	r.Handle("/contract/create", AuthMW(http.HandlerFunc(a.CreateContract))).Methods("POST")
	r.Handle("/contract/delete/{id}", AuthMW(http.HandlerFunc(a.DeleteContract))).Methods("DELETE")
	r.Handle("/contract/delete-all", AuthMW(http.HandlerFunc(a.DeleteAllContract))).Methods("POST")
	r.Handle("/contract/update/{id}", AuthMW(http.HandlerFunc(a.UpdateContract))).Methods("PUT")

	//BillControllers
	r.Handle("/bill/get-bills/{id}", AuthMW(http.HandlerFunc(a.GetBillsbyblock))).Methods("GET")
	r.Handle("/bill/get-bill-by-id/{id}", AuthMW(http.HandlerFunc(a.GetBills))).Methods("GET")
	r.Handle("/bill/create", AuthMW(http.HandlerFunc(a.CreateBill))).Methods("POST")
	r.Handle("/bill/delete/{id}", AuthMW(http.HandlerFunc(a.DeleteBill))).Methods("DELETE")
	r.Handle("/bill/update/{id}", AuthMW(http.HandlerFunc(a.UpdateBill))).Methods("PUT")
}
