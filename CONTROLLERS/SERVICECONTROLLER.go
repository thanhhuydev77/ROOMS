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

//get service of a bock with id from query
func (a *ApiDB) GetService(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	keys, ok := r.URL.Query()["idBlock"]
	if !ok || len(keys[0]) < 1 {
		io.WriteString(w, "{Url Param 'idBlock' is missing")
		return
	}
	idBlock, _ := strconv.Atoi(keys[0])
	listService, Ok := BUSINESS.GetServiceById(a.Db, idBlock)
	jsonlist, _ := json.Marshal(listService)
	if !Ok {
		io.WriteString(w, `{ "message": "Can’t get services" }`)
		return
	}
	stringresult := `{"status": 200,
    				"message": "Get services success",
    				"data": {
					"services":`
	if len(listService) > 0 {
		stringresult += string(jsonlist)
	} else {
		stringresult += "[]"
	}

	// stringresult += string(jsonlist)
	stringresult += "}}"
	io.WriteString(w, stringresult)
}

//delete a service with its id from variable
func (a *ApiDB) DeleteService(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	idservice, err := strconv.Atoi(vars["id"])

	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert id as int"}`)
		return
	}

	res, _ := BUSINESS.DeleteService(a.Db, idservice)

	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Delete service success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t  delete service"}`)
}

//create a service with information from body request
func (a *ApiDB) CreateService(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var p = MODELS.SERVICES_INPUT{}

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		io.WriteString(w, `{ "message": "Wrong format" }`)
		return
	}

	res, _ := BUSINESS.CreateService(a.Db, p.Services)
	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Create Services success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t create Services"}`)
}

//delete many services with id from body request
func (a *ApiDB) DeleteServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	p := struct {
		ServicesId []int `json:"servicesId"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		io.WriteString(w, `{ "message": "Wrong format" }`)
		return
	}
	res, _ := BUSINESS.DeleteServices(a.Db, p.ServicesId)

	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Delete Services success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t delete services"}`)
}

//update a service with information from body request
func (a *ApiDB) UpdateService(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	idService, _ := strconv.Atoi(vars["id"])

	var p = MODELS.SERVICE_INPUT{}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		io.WriteString(w, `{ "message": "Wrong format" }`)
		return
	}

	p.Id = idService
	res, _ := BUSINESS.UpdateService(a.Db, p)

	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Update service success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t update services"}`)
}
