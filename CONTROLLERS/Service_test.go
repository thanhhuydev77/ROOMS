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

func createMockService() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	result := sqlmock.NewRows([]string{"id", "nameService", "price", "idUnit", "description", "idBlock", "nameUnit"}).
		AddRow(119, "internet", 20000, 5, "tiền internet", 59, "GB").
		AddRow(120, "điện", 2000, 5, "tiền điện", 59, "kWh")
	mock.ExpectQuery(`SELECT S.*, U.name nameUnit FROM SERVICES S INNER JOIN UNITS U ON S.idUnit = U.id WHERE idBlock = \?`).WillReturnRows(result)

	return db, mock, err
}

func TestGetServicesPass(t *testing.T) {
	type data struct {
		Services []MODELS.GET_SERVICES_REQUEST `json:"services"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/service/get-services", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("idBlock", "59")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockService()
	a := &ApiDB{
		db,
	}

	handle := http.HandlerFunc(a.GetService)
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

	if out.Message != "Get services success" || len(out.Data.Services) == 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Services))
	}
}

func TestGetServicesFail1(t *testing.T) {
	type data struct {
		Services []MODELS.GET_SERVICES_REQUEST `json:"services"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/service/get-services", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockService()
	a := &ApiDB{
		db,
	}

	handle := http.HandlerFunc(a.GetService)
	handle.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetServicesFail2(t *testing.T) {
	type data struct {
		Services []MODELS.GET_SERVICES_REQUEST `json:"services"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/service/get-services", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("idBlock", "59")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	// db, _, _ := createMockService()
	a := &ApiDB{
		nil,
	}

	handle := http.HandlerFunc(a.GetService)
	handle.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func createMockDeleteService() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	result := sqlmock.NewRows([]string{"id", "nameService", "price", "idUnit", "description", "idBlock"}).
		AddRow(119, "internet", 20000, 5, "tiền internet", 59).
		AddRow(120, "điện", 2000, 5, "tiền điện", 59)
	mock.ExpectQuery(`delete from SERVICES`).WillReturnRows(result)

	return db, mock, err
}

func TestDeleteServicePass(t *testing.T) {
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("DELETE", "/service/delete/119", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "119",
	})

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeleteService()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.DeleteService)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var out output
	err = json.Unmarshal(rr.Body.Bytes(), &out)
	if err != nil {
		t.Errorf("error marshal :%v", err)
	}
	if out.Message != "Delete service success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func TestDeleteServiceFail1(t *testing.T) {
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("DELETE", "/service/delete/119", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "abc",
	})

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeleteService()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.DeleteService)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteServiceFail2(t *testing.T) {
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("DELETE", "/service/delete/119", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "119",
	})

	rr := httptest.NewRecorder()
	// db, _, _ := createMockDeleteService()
	a := &ApiDB{
		nil,
	}

	handler := http.HandlerFunc(a.DeleteService)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func createMockCreateService() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	result := sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2).AddRow(3)
	mock.ExpectQuery(`insert into SERVICES`).WillReturnRows(result)

	return db, mock, err
}

func TestCreateServicePass(t *testing.T) {

	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	var jsonStr = []byte(`
		{
			"services": [
				{	
					"nameService": "Tiền điện",
					"price": 3000,
					"idUnit": 4,
					"description": "Tiền điện",
					"idBlock": 51
				},
				{
					"nameService": "Tiền nước",
					"price": 22000,
					"idUnit": 5,
					"description": "Tiền nước",
					"idBlock": 51
				}
			]
		}
	`)

	req, err := http.NewRequest("POST", "/service/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockCreateService()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.CreateService)
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
	if out.Message != "Create Services success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func TestCreateServiceFail1(t *testing.T) {

	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	var jsonStr = []byte(`
		{
			abc
		}
	`)

	req, err := http.NewRequest("POST", "/service/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockCreateService()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.CreateService)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateServiceFail2(t *testing.T) {

	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	var jsonStr = []byte(`
		{
			"services": [
				{	
					"nameService": "Tiền điện",
					"price": 3000,
					"idUnit": 4,
					"description": "Tiền điện",
					"idBlock": 51
				},
				{
					"nameService": "Tiền nước",
					"price": 22000,
					"idUnit": 5,
					"description": "Tiền nước",
					"idBlock": 51
				}
			]
		}
	`)

	req, err := http.NewRequest("POST", "/service/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	// db, _, _ := createMockCreateService()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.CreateService)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func createMockDeleteServices() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	result := sqlmock.NewRows([]string{"id", "nameService", "price", "idUnit", "description", "idBlock"}).
		AddRow(119, "internet", 20000, 5, "tiền internet", 59).
		AddRow(120, "điện", 2000, 5, "tiền điện", 59)
	mock.ExpectQuery(`delete from SERVICES where`).WillReturnRows(result)

	return db, mock, err
}

func TestDeleteServicesPass(t *testing.T) {
	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	var jsonStr = []byte(`{
		"servicesId": [119,120]
	}`)

	req, err := http.NewRequest("POST", "/service/delete-all", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeleteServices()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.DeleteServices)
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
	if out.Message != "Delete Services success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func TestDeleteServicesFail1(t *testing.T) {
	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	var jsonStr = []byte(`{
		"servicesId": "[119,120]"
	}`)

	req, err := http.NewRequest("POST", "/service/delete-all", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeleteServices()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.DeleteServices)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var out Result
	err = json.Unmarshal(rr.Body.Bytes(), &out)
	if err != nil {
		t.Error("Can't run")
	}
}

func TestDeleteServicesFail2(t *testing.T) {
	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	var jsonStr = []byte(`{
		"servicesId": [119,120]
	}`)

	req, err := http.NewRequest("POST", "/service/delete-all", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	// db, _, _ := createMockDeleteServices()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.DeleteServices)
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
	if out.Message == "Delete Services success" {
		t.Errorf("Can't delete but: (%v)", out.Message)
	}
}
