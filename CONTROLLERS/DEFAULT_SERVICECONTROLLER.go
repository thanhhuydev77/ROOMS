package CONTROLLERS

import (
	"ROOMS/BUSINESS"
	"encoding/json"
	"io"
	"net/http"
)

//get a default service
func (a *ApiDB) Get_default_service(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	listUnit, Ok := BUSINESS.Get_all_default_service(a.Db)
	jsonlist, _ := json.Marshal(listUnit)
	if !Ok {
		io.WriteString(w, `{ "message": "Can’t get units" }`)
		return
	}
	stringresult := `{"status": 200,
    				"message": "Get default services success",
    				"data": {
        			"defaultServices":`
	stringresult += string(jsonlist)
	stringresult += "}}"
	io.WriteString(w, stringresult)
}
