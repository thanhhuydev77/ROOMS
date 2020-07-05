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

//get all block of user with id from query
func (a *ApiDB) GetBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	keys, ok := r.URL.Query()["userId"]
	if !ok || len(keys[0]) < 1 {
		io.WriteString(w, `{ "message": "Url Param 'userid' is missing"}`)
		return
	}
	idowner, _ := strconv.Atoi(keys[0])
	listBlock := BUSINESS.GetBlockByIdOwner(a.Db, idowner)
	jsonlist, _ := json.Marshal(listBlock)
	if len(listBlock) == 0 {
		io.WriteString(w, `{"status": 200,
    				"message": "Get Blocks success",
    				"data": {
        			"blocks": [] }}`)
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

//get a block with its id
func (a *ApiDB) GetBlockById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert idowner as int"}`)
		return
	}
	listBlock, _ := BUSINESS.GetBlockById(a.Db, id)
	jsonlist, _ := json.Marshal(listBlock)
	// if !OK {
	// 	io.WriteString(w, `{ "message": "Can’t get Blocks" }`)
	// 	return
	// }
	stringresult := `{"status": 200,
    				"message": "Get Blocks success",
    				"data": {
        			"blocks":`
	stringresult += string(jsonlist)
	stringresult += "}}"
	io.WriteString(w, stringresult)
}

//create a block with information from body request
func (a *ApiDB) CreateBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	p := MODELS.BLOCKS{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		io.WriteString(w, `{"message": "wrong format!"}`+err.Error())
		return
	}

	result, _ := BUSINESS.CreateBlock(a.Db, p)
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

//update a block with information from body request
func (a *ApiDB) UpdateBlock(w http.ResponseWriter, r *http.Request) {
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
		io.WriteString(w, `{"message": "wrong format!"}`)
		return
	}
	p.Id = idblock
	hasroweffected, _ := BUSINESS.UpdateBlock(a.Db, p)
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

//delete a block with id from variable
func (a *ApiDB) DeleteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	vars := mux.Vars(r)
	idblock, err := strconv.Atoi(vars["id"])

	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert id as int"}`)

		return
	}

	res, _ := BUSINESS.DeleteBlock(a.Db, idblock)

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

//delete many block with id from body request
func (a *ApiDB) DeleteBlocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var p = MODELS.IDBLOCKS{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		io.WriteString(w, `{ "message": "Wrong format" }`)
		return
	}
	res, _ := BUSINESS.DeleteBlocks(a.Db, p.BlocksId)

	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Delete Blocks success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t  Delete Block"}`)
}
