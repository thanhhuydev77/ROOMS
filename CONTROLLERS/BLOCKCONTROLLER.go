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

func GetBlockByOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	idowner, err := strconv.Atoi(vars["idowner"])
	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert idowner as int"}`)
		return
	}
	listBlock := BUSINESS.GetBlockByIdOwner(idowner)
	jsonlist, _ := json.Marshal(listBlock)
	if len(listBlock) == 0 {
		io.WriteString(w, `{ "message": "Can’t get Blocks" }`)
		return
	}
	stringresult := `{"status": 200,
    				"message": "Get Blocks success",
    				"data": {
        			"blocks":`
	stringresult += string(jsonlist)
	stringresult += "}}"
	io.WriteString(w, stringresult)
}

func CreateBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	p := MODELS.BLOCKS{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		io.WriteString(w, `{"message": "wrong format!"}`+err.Error())
		return
	}

	result, _ := BUSINESS.CreateBlock(p)
	if result {
		io.WriteString(w, `{ "status": 200,
    						"message": "Create block success",
    							"data": {
        						"status": 1
    									}
								}`)
	} else {
		io.WriteString(w, `{ "message": "Can’t create block "}`)
	}
}

func UpdateBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	idblock, err := strconv.Atoi(vars["Idblock"])
	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert id as int"}`)
		return
	}
	p := MODELS.BLOCKS{}
	err1 := json.NewDecoder(r.Body).Decode(&p)
	if err1 != nil {
		io.WriteString(w, `{"message": "wrong format!"}`+err.Error())
		return
	}
	p.Id = idblock
	hasroweffected, _ := BUSINESS.UpdateBlock(p)
	if hasroweffected == false {
		io.WriteString(w, `{ "message": "Can’t update block" }`)
		return
	}
	stringresult := `{  "status": 200,
    					"message": "Update Success",
						"data": {
        						"status": 1
    							}
						}`
	io.WriteString(w, stringresult)
	return
}
