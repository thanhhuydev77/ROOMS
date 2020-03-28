package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	//r.HandleFunc("/Bills/getall/{id}", func(w http.ResponseWriter, r *http.Request) {
	//	vars := mux.Vars(r)
	//	_,err := strconv.Atoi(vars["id"])
	//	bill := MODELS.BILLS{}
	//	bill.IdRoom = r.GetBody
	//	if err == nil {
	//		json.NewEncoder(w).Encode(BUSINESS.CreateBill())
	//	}
	//
	//})
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", r)
}
