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

func CreateRoom(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("Content-Type", "application/json")

	room := MODELS.ROOMS{}
	err := json.NewDecoder(r.Body).Decode(&room)

	if err != nil{
		io.WriteString(w, `{"message": "Wrong format!"}`)
		return
	}

	result, err := BUSINESS.CreateRoom(room)
	if result{
		io.WriteString(w, `{
						"status": 200,
						"message": "Create rooms success",
						"data": {
							"status": 1
							}
						}`)
	}else{
		io.WriteString(w,`{"message" : "Can't create room'"}`)
	}
}

func DeleteRoom(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil{
		io.WriteString(w , `"message": "Can’t delete room"`)
	}

	result, err := BUSINESS.DeleteRoom(id)

	if result{
		io.WriteString(w , ` "status": 200,
								"message": "Delete rooms success",
								"data": {
									"status": 1
								}`)
	}else{
		io.WriteString(w , `"message": "Can’t delete room"`)
	}
}