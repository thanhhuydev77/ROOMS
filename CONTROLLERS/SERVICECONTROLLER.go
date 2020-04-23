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

func GetService(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	keys, ok := r.URL.Query()["idBlock"]
	if !ok || len(keys[0]) < 1 {
		io.WriteString(w, "{Url Param 'idBlock' is missing")
		return
	}
	idBlock, _ := strconv.Atoi(keys[0])
	listBlock, Ok := BUSINESS.GetServiceById(idBlock)
	jsonlist, _ := json.Marshal(listBlock)
	if !Ok {
		io.WriteString(w, `{ "message": "Can’t get services" }`)
		return
	}
	stringresult := `{"status": 200,
    				"message": "Get services success",
    				"data": {
        			"services":`
	stringresult += string(jsonlist)
	stringresult += "}}"
	io.WriteString(w, stringresult)
}

func DeleteService(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	idservice, err := strconv.Atoi(vars["id"])

	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert id as int"}`)
		return
	}

	res, _ := BUSINESS.DeleteService(idservice)

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

func CreateService(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("Content-Type", "application/json")

	var p = MODELS.SERVICES_INPUT{}

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		io.WriteString(w, `{ "message": "Wrong format" }`)
		return
	}

	res, _ := BUSINESS.CreateService(p.Services)
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
