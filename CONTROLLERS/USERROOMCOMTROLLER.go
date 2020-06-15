package CONTROLLERS

import (
	"ROOMS/BUSINESS"
	"encoding/json"
	"io"
	"net/http"
)

func DeleteManyUserRoom(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("Content-Type", "application/json")

	ids := struct {
		UserRoomIds []int
	}{}

	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		io.WriteString(w, `{ "message": "Wrong format" }` + err.Error())
		return
	}

	result, err := BUSINESS.DeleteManyUserRoom(ids.UserRoomIds)

	if result {
		io.WriteString(w, `{
						"status": 200,
						"message": "Delete user room success",
						"data": {
							"status": 1
							}
						}`)
		return
	} else {
		io.WriteString(w, `{"message" : "Can’t delete user room"}`)
	}
}