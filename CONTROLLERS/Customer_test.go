package CONTROLLERS

import (
	"ROOMS/MODELS"
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
)

func createMockCustomer() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	result := sqlmock.NewRows([]string{"id", "codeUser", "userName", "pass", "fullName", "identifyFront", "identifyBack",
		"dateBirth", "address", "role", "sex", "job", "workPlace", "tempReg", "province", "email", "avatar",
		"phoneNumber", "idOwner", "note"}).
		AddRow(111, "#1589559118", "null", "null", "Lâm Khắc Duy", "https://vi.api.vinlt.wtf/public/images/avatars/1589559112428-1.PNG",
			"https://vi.api.vinlt.wtf/public/images/avatars/1589559116076-1.PNG", nil, "null", 1, "male", "Sinh Viên",
			"TP HCM", 1, "null", "khacduylam@gmail.com", "https://vi.api.vinlt.wtf/public/images/avatars/1589559108990-1.PNG",
			"03425251111", 69, "e21e21 rewrwerew fwefwe")
	mock.ExpectQuery(`SELECT \* FROM CUSTOMERS WHERE idOwner = ?`).WillReturnRows(result)
	mock.ExpectQuery(`SELECT R.nameRoom FROM USER_ROOM UR INNER JOIN ROOMS R ON UR.idRoom = R.id WHERE idUser = \?`).WillReturnRows(result)
	return db, mock, err
}

func TestGetCustomersPass(t *testing.T) {
	type data struct {
		Customers []MODELS.CUSTOMER_GET `json:"customers"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/customer/get-customers", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("userId", "69")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockCustomer()
	a := &ApiDB{
		db,
	}

	handle := http.HandlerFunc(a.GetCustomersByUserId)
	handle.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var out output
	err = json.Unmarshal(rr.Body.Bytes(), &out)
	if err != nil {
		t.Errorf("error marshal :%v", err)
	}

	if out.Message != "Get customers success" || len(out.Data.Customers) == 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Customers))
	}
}

func TestGetCustomersFail1(t *testing.T) {
	type data struct {
		Customers []MODELS.CUSTOMER_GET `json:"customers"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/customer/get-customers", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("userId", "")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockCustomer()
	a := &ApiDB{
		db,
	}

	handle := http.HandlerFunc(a.GetCustomersByUserId)
	handle.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetCustomersFail2(t *testing.T) {
	type data struct {
		Customers []MODELS.CUSTOMER_GET `json:"customers"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/customer/get-customers", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("userId", "abc")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	a := &ApiDB{
		nil,
	}

	handle := http.HandlerFunc(a.GetCustomersByUserId)
	handle.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func createMockCreateCustomer() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	result := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("INSERT INTO CUSTOMERS").WillReturnRows(result)

	return db, mock, err
}

func TestCreateCustomerPass(t *testing.T) {

	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	var jsonStr = []byte(`{
		"fullName": "Nguyễn Thành Huy 2",
		"phoneNumber": "0987898767",
		"dateBirth": "2020-05-11",
		"email": "nltruongvi@gmail.com",
		"job": "Người đi làm",
		"workPlace": "Trường đại học công nghệ thông tin",
		"sex": "male",
		"tempReg": 1,
		"note": "fqw2132121",
		"avatar": "",
		"identifyBack": "",
		"identifyFront": "",
		"codeUser": "#1589206969",
		"idOwner": 40
		}`)

	req, err := http.NewRequest("POST", "/customer/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockCreateCustomer()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.CreateCustomer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var out Result
	err = json.Unmarshal(rr.Body.Bytes(), &out)
	if err != nil {
		t.Errorf("error marshal :%v", err)
	}
	if out.Message != "Create customer success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func TestCreateCustomerFail1(t *testing.T) {

	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	var jsonStr = []byte(`{
			abc
		}`)

	req, err := http.NewRequest("POST", "/customer/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockCreateCustomer()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.CreateCustomer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateCustomerFail2(t *testing.T) {

	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	var jsonStr = []byte(`{
		"fullName": "Nguyễn Thành Huy 2",
		"phoneNumber": "0987898767",
		"dateBirth": "2020-05-11",
		"email": "nltruongvi@gmail.com",
		"job": "Người đi làm",
		"workPlace": "Trường đại học công nghệ thông tin",
		"sex": "male",
		"tempReg": 1,
		"note": "fqw2132121",
		"avatar": "",
		"identifyBack": "",
		"identifyFront": "",
		"codeUser": "#1589206969",
		"idOwner": 40
		}`)

	req, err := http.NewRequest("POST", "/customer/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.CreateCustomer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func createMockDeleteCustomer() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	result := sqlmock.NewRows([]string{"id", "codeUser", "userName", "pass", "fullName", "identifyFront", "identifyBack",
		"dateBirth", "address", "role", "sex", "job", "workPlace", "tempReg", "province", "email", "avatar",
		"phoneNumber", "idOwner", "note"}).
		AddRow(111, "#1589559118", "null", "null", "Lâm Khắc Duy", "https://vi.api.vinlt.wtf/public/images/avatars/1589559112428-1.PNG",
			"https://vi.api.vinlt.wtf/public/images/avatars/1589559116076-1.PNG", nil, "null", 1, "male", "Sinh Viên",
			"TP HCM", 1, "null", "khacduylam@gmail.com", "https://vi.api.vinlt.wtf/public/images/avatars/1589559108990-1.PNG",
			"03425251111", 69, "e21e21 rewrwerew fwefwe")
	mock.ExpectQuery(`DELETE FROM CUSTOMERS`).WillReturnRows(result)
	// mock.ExpectQuery(`SELECT R.nameRoom FROM USER_ROOM UR INNER JOIN ROOMS R ON UR.idRoom = R.id WHERE idUser = \?`).WillReturnRows(result)
	return db, mock, err
}

func TestDeleteCustomerPass(t *testing.T) {
	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("delete", "/customer/delete/111", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "68",
	})

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeleteCustomer()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.DeleteCustomer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var out Result
	err = json.Unmarshal(rr.Body.Bytes(), &out)
	if err != nil {
		t.Errorf("error marshal :%v", err)
	}
	if out.Message != "Delete customer success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func TestDeleteCustomerFail1(t *testing.T) {
	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("delete", "/customer/delete/111", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "",
	})

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeleteCustomer()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.DeleteCustomer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteCustomerFail2(t *testing.T) {
	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("delete", "/customer/delete/111", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "111",
	})

	rr := httptest.NewRecorder()
	// db, _, _ := createMockDeleteCustomer()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.DeleteCustomer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
