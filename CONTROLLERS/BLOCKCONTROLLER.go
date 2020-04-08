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

func GetBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	keys, ok := r.URL.Query()["userId"]
	if !ok || len(keys[0]) < 1 {
		io.WriteString(w, "{Url Param 'userid' is missing")
		return
	}
	idowner, _ := strconv.Atoi(keys[0])
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

func GetBlockBYId(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert idowner as int"}`)
		return
	}
	listBlock, OK := BUSINESS.GetBlockById(id)
	jsonlist, _ := json.Marshal(listBlock)
	if !OK {
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
	idblock, err := strconv.Atoi(vars["id"])
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

func DeleteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	idblock, err := strconv.Atoi(vars["id"])

	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert id as int"}`)
		return
	}

	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert id as int"}`)
		return
	}

	res, _ := BUSINESS.DeleteBlock(idblock)

	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Delete Block success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t  Delete Block"}`)
}

func DeleteBlocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var p = MODELS.IDBLOCKS{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		io.WriteString(w, `{ "message": "Wrong format" }`)
		return
	}
	res, _ := BUSINESS.DeleteBlocks(p.BlocksId)

	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Delete Block success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t  Delete Block"}`)
}
