package CONTROLLERS

import (
	"ROOMS/BUSINESS"
	"ROOMS/MODELS"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	idBlockP, ok := r.URL.Query()["idBlock"]
	if !ok || len(idBlockP[0]) < 1 {
		io.WriteString(w, `{ "message": "Url Param 'idBlock' is missing"}`)
		return
	}
	idBlock, err := strconv.Atoi(idBlockP[0])
	if err != nil {
		io.WriteString(w, `{"message": "Wrong format!"}`)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %s", err)
		return
	}
	isMatchP, err := strconv.ParseBool(r.Form.Get("isMatch"))

	if isMatchP {
		result, bool, _ := BUSINESS.UpdateGetRoom(idBlock)
		resultJson, _ := json.Marshal(result)

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
	} else {
		result, bool, _ := BUSINESS.GetRoom(idBlock)
		resultJson, _ := json.Marshal(result)

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

func GetRoomDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	status, ok := r.URL.Query()["status"]
	idBlock, ok1 := r.URL.Query()["idBlock"]
	userId, ok2 := r.URL.Query()["userId"]
	if !ok1 || len(idBlock[0]) < 1 {
		io.WriteString(w, `{ "message": "Url Param 'idBlock' is missing"}`)
		return
	}
	if !ok || len(status[0]) < 1 {
		io.WriteString(w, `{ "message": "Url Param 'status' is missing"}`)
		return
	}
	if !ok2 || len(userId[0]) < 1 {
		io.WriteString(w, `{ "message": "Url Param 'userId' is missing"}`)
		return
	}

	IdBlock, err := strconv.Atoi(idBlock[0])
	Status, err1 := strconv.Atoi(status[0])
	UserId, err2 := strconv.Atoi(userId[0])

	if err != nil || err1 != nil || err2 != nil {
		io.WriteString(w, `{"message": "Wrong format!"}`)
		return
	}

	result, err := BUSINESS.GetRoomDB(IdBlock, Status, UserId)
	resultJson, _ := json.Marshal(result)
	if err != nil {
		io.WriteString(w, `{"message": "Can’t get rooms"}`)
	}
	data := `{		"status": 200,
					"message": "Get rooms success",
					"data":
						{"rooms":`
	data += string(resultJson)
	data += `}}`
	io.WriteString(w, data)

}

func GetRoomImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	codeRoom, ok := r.URL.Query()["codeRoom"]
	if !ok || len(codeRoom[0]) < 1 {
		io.WriteString(w, `{ "message": "Url Param 'idBlock' is missing"}`)
		return
	}

	result, bool, _ := BUSINESS.GetRoomImage(codeRoom[0])
	resultJson, _ := json.Marshal(result)
	var data string
	if bool {
		data = `{		"status": 200,
					"message": "Get images success",
					"data":
						{"rooms":`
		if len(result) > 0 {
			data += string(resultJson)
		}
		data += `}}`

	} else {
		data = `{    "message": "Can’t get images",}`

	}
	io.WriteString(w, data)

}
