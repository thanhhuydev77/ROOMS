package CONTROLLERS

import (
	"ROOMS/BUSINESS"
	"ROOMS/MODELS"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func GetCustomersByUserId(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("Content-Type","application/json")
	keys, ok := r.URL.Query()["userId"]
	if !ok || len(keys[0]) < 1 {
		io.WriteString(w, `{"message":"can not convert idowner as int"}`)
		return
	}
	userId, _ := strconv.Atoi(keys[0])
	listCustomer, _, err := BUSINESS.GetCustomersByUserId(userId)
	jsonCustomers, _ := json.Marshal(listCustomer)

	if err != nil{
		io.WriteString(w, `{ "message": "Can’t get customers" }`)
		return
	}

	stringresult := `{"status": 200,
    				"message": "Get customers success",
    				"data": {
        			"customers":`
	stringresult += string(jsonCustomers)
	stringresult += "}}"
	io.WriteString(w, stringresult)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("Content-Type","application/json")

	p := MODELS.CUSTOMER_INPUT{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil{
		io.WriteString(w, `{"message": "wrong format!"}` + err.Error())
		return
	}

	rs, _ := BUSINESS.CreateCustomer(p)

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