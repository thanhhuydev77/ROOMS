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

func (a *ApiDB) GetBills(w http.ResponseWriter, r *http.Request) {

	type Data struct {
		Bill       MODELS.BILLS          `json:"bill"`
		Billdetail []MODELS.BILL_DETAILS `json:"billdetail"`
	}
	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}

	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	idBill, err := strconv.Atoi(vars["id"])

	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert id as int"}`)

		return
	}
	billinfo, okinfo, err := BUSINESS.GetBillById(a.Db, idBill)
	billdetail, okdetail, err := BUSINESS.GetBillDetailById(a.Db, idBill)
	if okinfo != true || okdetail != true {
		io.WriteString(w, `{ "message": "Can’t get contracts" }`)
		return
	}
	result := Result{
		Status:  200,
		Message: "get bill success",
		Data: Data{
			Bill:       billinfo,
			Billdetail: billdetail,
		},
	}

	jsonresult, _ := json.Marshal(result)
	io.WriteString(w, string(jsonresult))
}

func (a *ApiDB) CreateBill(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	p := MODELS.CREATE_UPDATE_BILL_REQUEST{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		io.WriteString(w, `{"message": "wrong format!"}`)
		return
	}

	result, _ := BUSINESS.CreateBill(a.Db, p)

	if result > 0 {
		io.WriteString(w, `  { "status": 200,
    "message": "Create bill success",
    "data": {
        "status": 1
    }
}
`)
	} else {

		io.WriteString(w, `{ "message": "Can’t create bill"}`)
	}
}

func (a *ApiDB) UpdateBill(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	idContract, _ := strconv.Atoi(vars["id"])

	var p = MODELS.CREATE_UPDATE_BILL_REQUEST{}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		io.WriteString(w, `{ "message": "Wrong format" }`)
		return
	}

	p.Id = idContract
	res, _ := BUSINESS.UpdateBill(a.Db, p)

	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Update bill success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t update bill"}`)
}

func (a *ApiDB) DeleteBill(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	idbill, err := strconv.Atoi(vars["id"])

	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"can not convert id as int"}`)

		return
	}

	res, _ := BUSINESS.DeleteBill(a.Db, idbill)

	if res {
		io.WriteString(w, `{
						"status": 200,
						"message": "Delete bill success",
						"data": {
							"status": 1
							}
						}`)
		return
	}
	io.WriteString(w, `{"message" : "Can’t delete bill"}`)
}

func (a *ApiDB) GetBillsbyblock(w http.ResponseWriter, r *http.Request) {

	type Data struct {
		Bills []MODELS.BILLS `json:"bills"`
	}

	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	idBlock, err := strconv.Atoi(vars["id"])

	if err != nil {
		io.WriteString(w, `{"message":"can not convert id as int"}`)
		return
	}
	bill, ok, errget := BUSINESS.GetBillByIdblock(a.Db, idBlock)
	if ok != true || errget != nil {
		//fmt.Print(errget)
		io.WriteString(w, `{ "message": "Can’t get bills" }`)
		return
	}

	res := MODELS.RespondOk{
		Status:  200,
		Message: "get bill success",
		Data:    Data{bill},
	}
	jsonresult, _ := json.Marshal(res)

	io.WriteString(w, string(jsonresult))
}
