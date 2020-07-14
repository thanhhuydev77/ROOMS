package CONTROLLERS

import (
	"ROOMS/BUSINESS"
	"ROOMS/MODELS"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

//get all contract with blockid from query
func (a *ApiDB) GetContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	keys, ok := r.URL.Query()["idBlock"]
	if !ok || len(keys[0]) < 1 {
		io.WriteString(w, `{"message":"can not convert idBlock as int"}`)
		return
	}
	BlockId, _ := strconv.Atoi(keys[0])
	listCustomer, _, err := BUSINESS.GetContractByBlockId(a.Db, BlockId)
	jsonCustomers, _ := json.Marshal(listCustomer)

	if err != nil {
		io.WriteString(w, `{ "message": "Can’t get contracts" }`)
		return
	}

	stringresult := `{"status": 200,
    				"message": "Get contracts success",
    				"data": {
        			"contracts":`
	stringresult += string(jsonCustomers)
	stringresult += "}}"
	io.WriteString(w, stringresult)
}

//create a contract with information from body request
func (a *ApiDB) CreateContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	p := MODELS.CREATE_UPDATE_CONTRACT_REQUEST{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		io.WriteString(w, `{"message": "wrong format!"}`)
		return
	}

	result := BUSINESS.CreateContract(a.Db, p)
	if result {
		io.WriteString(w, `  { "status": 200,
    "message": "Create contract success",
    "data": {
        "status": 1
    }
}
`)
	} else {
		io.WriteString(w, `{ "message": "Can’t create contract"}`)
	}
}

//delete a contract with id from variable
func (a *ApiDB) DeleteContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	idContract, err := strconv.Atoi(vars["id"])

	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert id as int"}`)

		return
	}

	res, _ := BUSINESS.DeleteContract(a.Db, idContract)

	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Delete contract success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t delete contract"}`)
}

//delete many contract with id from body request
func (a *ApiDB) DeleteAllContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	type arrcontractId struct {
		ContractsId []int `json:"contractsId"`
	}
	p := arrcontractId{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		io.WriteString(w, `{"message": "wrong format!"}`)
		return
	}

	res, _ := BUSINESS.DeleteAllContract(a.Db, p.ContractsId)

	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Delete contract success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t delete contract"}`)
}

//update a contract with information from body request
func (a *ApiDB) UpdateContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	idContract, _ := strconv.Atoi(vars["id"])

	var p = MODELS.CREATE_UPDATE_CONTRACT_REQUEST{}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		io.WriteString(w, `{ "message": "Wrong format" }`)
		return
	}

	p.Id = idContract
	res, _ := BUSINESS.UpdateContract(a.Db, p)

	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Update contracts success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t update contracts"}`)
}
