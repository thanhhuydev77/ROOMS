package CONTROLLERS

import (
	"ROOMS/BUSINESS"
	"encoding/json"
	"io"
	"net/http"
)

func (a *ApiDB) GetAllUnit(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	listUnit, Ok := BUSINESS.GetAllUnits(a.Db)
	jsonlist, _ := json.Marshal(listUnit)
	if !Ok {
		io.WriteString(w, `{ "message": "Can’t get units" }`)
		return
	}
	stringresult := `{"status": 200,
    				"message": "Get units success",
    				"data": {
        			"units":`
	stringresult += string(jsonlist)
	stringresult += "}}"
	io.WriteString(w, stringresult)
}
