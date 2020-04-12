package CONTROLLERS

import (
	"ROOMS/BUSINESS"
	"ROOMS/MODELS"
	"encoding/json"
	"io"
	"net/http"
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