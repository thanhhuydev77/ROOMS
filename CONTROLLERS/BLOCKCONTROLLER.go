package CONTROLLERS

import (
	"ROOMS/BUSINESS"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func GetBlockByOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idowner, err := strconv.Atoi(vars["idowner"])
	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert idowner as int"}`)
		return
	}
	listBlock := BUSINESS.GetBlockByIdOwner(idowner)
	stringresult := `{"message": "get blocks success","data":{`
	for _, val := range listBlock {
		stringresult += `{"Id":"` + strconv.Itoa(val.Id) + `","Name":"` + val.Name + `","Address":"` + val.Address + `","Description":"` + val.Description + `"},`
	}
	stringresult += "}}"
	io.WriteString(w, stringresult)
}
