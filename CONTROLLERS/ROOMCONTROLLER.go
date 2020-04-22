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

func GetRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	keys, ok := r.URL.Query()["idBlock"]
	if !ok || len(keys[0]) < 1 {
		io.WriteString(w, "{Url Param 'idBlock' is missing")
		return
	}

	idBlock, err := strconv.Atoi(keys[0])

	if err != nil {
		io.WriteString(w, `{"message": "Wrong format!"}`)
		return
	}

	result, bool, _ := BUSINESS.GetRoom(idBlock)
	resultJson, err := json.Marshal(result)

	data := `{		"status": 200,
					"message": "Get rooms success",
					"data":
						{"rooms":`
	if len(result) > 0 {
		data += string(resultJson)
	}
	data += `}}`

	if bool {
		io.WriteString(w, data)
	}
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	room := MODELS.ROOMS{}
	err := json.NewDecoder(r.Body).Decode(&room)

	if err != nil {
		io.WriteString(w, `{"message": "Wrong format!"}`)
		return
	}

	result, err := BUSINESS.CreateRoom(room)
	if result {
		io.WriteString(w, `{
						"status": 200,
						"message": "Create rooms success",
						"data": {
							"status": 1
							}
						}`)
	} else {
		io.WriteString(w, `{"message" : "Can't create room'"}`)
	}
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		io.WriteString(w, `{"message": "Wrong format!"}`)
		return
	}

	result, err := BUSINESS.DeleteRoom(id)

	if result {
		io.WriteString(w, `{ 
								"status": 200,
								"message": "Delete rooms success",
								"data": {
									"status": 1
								}}`)
	} else {
		io.WriteString(w, `"message": "Can’t delete room"`)
	}
}

func DeleteRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var ids = MODELS.ROOMIDS{}
	err := json.NewDecoder(r.Body).Decode(&ids)

	if err != nil {
		io.WriteString(w, "Wrong format !")
		return
	}

	result, _ := BUSINESS.DeleteRooms(ids.RoomsId)

	if result {
		io.WriteString(w, `{
								"status": 200,
								"message": "Delete rooms success",
								"data": {
									"status": 1
								}
							}`)
	} else {
		io.WriteString(w, `{"message": "Can’t delete rooms"}`)
	}
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	room := MODELS.ROOMS{}
	arr2 := json.NewDecoder(r.Body).Decode(&room)

	if err != nil || arr2 != nil {
		io.WriteString(w, `{"message": "Wrong format!"}`)
		return
	}

	result, err := BUSINESS.UpdateRoom(id, room)

	if result {
		io.WriteString(w, `{  "status": 200,
								"message": "Update room success",
								"data": {
								"status": 1
								}}`)
	} else {
		io.WriteString(w, `{"message": "Can’t update room"}`)
	}
}
