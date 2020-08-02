package CONTROLLERS

import (
	"ROOMS/BUSINESS"
	"ROOMS/MODELS"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//get customer infoation from iduser from query
func (a *ApiDB) GetCustomersByUserId(w http.ResponseWriter, r *http.Request) {
	type paging struct {
		TotalRows int `json:"_totalRows"`
	}
	type Data struct {
		Page      int                   `json:"page"`
		Limit     int                   `json:"limit"`
		TotalRows int                   `json:"totalRows"`
		Customers []MODELS.CUSTOMER_GET `json:"customers"`
	}
	type Respond struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Dataa   Data   `json:"data"`
	}

	w.Header().Add("Content-Type", "application/json")
	keys, ok := r.URL.Query()["userId"]
	pagenum, ok := r.URL.Query()["page"]
	limits, ok := r.URL.Query()["limit"]
	if !ok || len(keys[0]) < 1 {
		io.WriteString(w, `{"message":"can not convert idowner as int"}`)
		return
	}
	if !ok || len(pagenum[0]) < 1 {
		io.WriteString(w, `{"message":"can not convert page as int"}`)
		return
	}
	if !ok || len(limits[0]) < 1 {
		io.WriteString(w, `{"message":"can not convert limit as int"}`)
		return
	}
	var Page paging

	userId, _ := strconv.Atoi(keys[0])
	page, _ := strconv.Atoi(pagenum[0])
	limit, _ := strconv.Atoi(limits[0])
	listCustomer, _, err, numrow := BUSINESS.GetCustomersByUserId(a.Db, userId, (page-1)*limit, limit)
	Page.TotalRows = numrow

	//jsonCustomers, _ := json.Marshal(listCustomer)

	if err != nil {
		io.WriteString(w, `{ "message": "Can’t get customers" }`)
		return
	}
	respond := Respond{
		Status:  200,
		Message: "Get customers success",
		Dataa: Data{
			Page:      page,
			Limit:     limit,
			TotalRows: numrow,
			Customers: listCustomer,
		},
	}
	jsonresult, _ := json.Marshal(respond)
	io.WriteString(w, string(jsonresult))
}

//create a customer with information from body request
func (a *ApiDB) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	p := MODELS.CUSTOMER_INPUT{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		io.WriteString(w, `{"message": "wrong format!"}`)
		return
	}

	rs, _ := BUSINESS.CreateCustomer(a.Db, p)

	if rs {
		io.WriteString(w, `{ "status": 200,
    						"message": "Create customer success",
    							"data": {
        						"status": 1
    									}
								}`)
	} else {
		io.WriteString(w, `{ "message": "Can’t create customer "}`)
	}
}

//delete a customer with its id from variable
func (a *ApiDB) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	vars := mux.Vars(r)
	idCustomer, err := strconv.Atoi(vars["id"])

	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert id as int"}`)

		return
	}

	res, _ := BUSINESS.DeleteCustomer(a.Db, idCustomer)

	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Delete customer success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t delete customer"}`)
}

//delete many customer with id from body request
func (a *ApiDB) DeleteManyCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var p = MODELS.CUSTOMERIDS{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		io.WriteString(w, `{ "message": "Wrong format" }`)
		return
	}
	res, _ := BUSINESS.DeleteManyCustomers(a.Db, p.CustomersId)

	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Delete customers success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t  delete customers"}`)
}

//update a customer with information from body request
func (a *ApiDB) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	idCustomer, err := strconv.Atoi(vars["id"])
	if err != nil {
		io.WriteString(w, `{"message":"can not convert id as int"}`)
		return
	}

	p := MODELS.CUSTOMER_UPDATE{}
	err1 := json.NewDecoder(r.Body).Decode(&p)
	if err1 != nil {
		io.WriteString(w, `{"message": "wrong format!"}`)
		return
	}
	p.Id = idCustomer
	hasroweffected, _ := BUSINESS.UpdateCustomer(a.Db, p)
	if hasroweffected == false {
		io.WriteString(w, `{ "message": "Can’t update customer" }`)
		return
	}
	stringresult := `{  "status": 200,
    					"message": "Update customer Success",
						"data": {
        						"status": 1
    							}
						}`
	io.WriteString(w, stringresult)
	return
}
