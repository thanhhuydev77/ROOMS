package CONTROLLERS

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createMockGetContract() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	DeleteRoom := sqlmock.NewRows([]string{"Id", "IdRoom", "IdOwner", "IdSlave", "StartDate",
		"EndDate", "CirclePay", "Deposit", "DayPay", "Note",
		"IdBlock", "NameRoom", "FullName", "IdUsers"}).
		AddRow(1, 1, 1, 1, "name1", "dwa", 1, float64(1), 1, "no descrip", 1, "dwa", "dwa", 1)
	mock.ExpectQuery(`SELECT C.*, R.nameRoom, CU.fullName FROM CONTRACTS`).WillReturnRows(DeleteRoom)
	mock.ExpectQuery(`SELECT idUser FROM USER_ROOM WHERE`).WillReturnRows(sqlmock.NewRows([]string{"iduser"}).AddRow(1))
	//mock.ExpectCommit()
	return db, mock, err
}
func TestApiDB_GetContractPass(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("idBlock", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	db, _, _ := createMockGetContract()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.GetContract)
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
	if out.Message != "Get contracts success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestApiDB_GetContractFail(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}

	//q := req.URL.Query()
	//q.Add("idBlock", "1")
	//req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	db, _, _ := createMockGetContract()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.GetContract)
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
	if out.Message == "Get contracts success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestApiDB_GetContractFail2(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("idBlock", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	//db, _, _ := createMockGetContract()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.GetContract)
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
	if out.Message == "Get contracts success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockCreateContract() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO CONTRACTS`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO USER_ROOM`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	return db, mock, err
}
func TestApiDB_CreateContractPass(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
				"id":1,
				"idRoom":1,
				"idOwner":1,
				"idSlave":1,
				"startDate":"abc",
				"endDate":"xyz",
				"circlePay":1,
				"deposit":1,
				"dayPay":1,
				"note":"abc",
				"idBlock":1,
				"userRooms":[
							{
								"idUser":1,
								"idRoom":1}
							]
}`)
	req, err := http.NewRequest("POST", "/room/create", bytes.NewBuffer(jsonStr))
	//req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("idBlock", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	db, _, _ := createMockCreateContract()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.CreateContract)
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
	if out.Message != "Create contract success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestApiDB_CreateContractFail(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{,,}`)
	req, err := http.NewRequest("POST", "/room/create", bytes.NewBuffer(jsonStr))
	//req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("idBlock", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	db, _, _ := createMockCreateContract()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.CreateContract)
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
	if out.Message == "Create contract success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestApiDB_CreateContractFail2(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
				"id":1,
				"idRoom":1,
				"idOwner":1,
				"idSlave":1,
				"startDate":"abc",
				"endDate":"xyz",
				"circlePay":1,
				"deposit":1,
				"dayPay":1,
				"note":"abc",
				"idBlock":1,
				"userRooms":[
							{
								"idUser":1,
								"idRoom":1}
							]
}`)
	req, err := http.NewRequest("POST", "/room/create", bytes.NewBuffer(jsonStr))
	//req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("idBlock", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	//db, _, _ := createMockCreateContract()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.CreateContract)
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
	if out.Message == "Create contract success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockDeleteContract() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	//mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM CONTRACTS`).WillReturnResult(sqlmock.NewResult(1, 1))
	//mock.ExpectExec(`INSERT INTO USER_ROOM`).WillReturnResult(sqlmock.NewResult(1,1))
	//mock.ExpectCommit()
	return db, mock, err
}
func TestApiDB_DeleteContractPass(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/room/create", nil)
	//req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeleteContract()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.DeleteContract)
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
	if out.Message != "Delete contract success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestApiDB_DeleteContractFail(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/room/create", nil)
	//req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}
	//req = mux.SetURLVars(req, map[string]string{
	//	"id": "1",
	//})

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeleteContract()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.DeleteContract)
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
	if out.Message == "Delete contract success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestApiDB_DeleteContractFail2(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/room/create", nil)
	//req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	//db, _, _ := createMockDeleteContract()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.DeleteContract)
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
	if out.Message == "Delete contract success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func TestApiDB_DeleteAllContractPass(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
		"contractsId":[1]
}`)
	req, err := http.NewRequest("POST", "/room/create", bytes.NewBuffer(jsonStr))
	//req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeleteContract()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.DeleteAllContract)
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
	if out.Message != "Delete contract success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestApiDB_DeleteAllContractFail(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
		"contractsId":[a]
}`)
	req, err := http.NewRequest("POST", "/room/create", bytes.NewBuffer(jsonStr))
	//req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeleteContract()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.DeleteAllContract)
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
	if out.Message == "Delete contract success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestApiDB_DeleteAllContractPassFail2(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
		"contractsId":[1]
}`)
	req, err := http.NewRequest("POST", "/room/create", bytes.NewBuffer(jsonStr))
	//req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	//db, _, _ := createMockDeleteContract()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.DeleteAllContract)
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
	if out.Message == "Delete contract success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockUpdateContract() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE CONTRACTS SET `).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`DELETE FROM USER_ROOM`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO USER_ROOM`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	return db, mock, err
}
func TestApiDB_UpdateContractPass(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
				"id":1,
				"idRoom":1,
				"idOwner":1,
				"idSlave":1,
				"startDate":"abc",
				"endDate":"xyz",
				"circlePay":1,
				"deposit":1,
				"dayPay":1,
				"note":"abc",
				"idBlock":1,
				"userRooms":[
							{
								"idUser":1,
								"idRoom":1}
							]
}`)
	req, err := http.NewRequest("POST", "/room/create", bytes.NewBuffer(jsonStr))
	//req, err := http.NewRequest("POST", "/room/create", nil)
	//req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	db, _, _ := createMockUpdateContract()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.UpdateContract)
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
	if out.Message != "Update contracts success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestApiDB_UpdateContractFail(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{,,,}`)
	req, err := http.NewRequest("POST", "/room/create", bytes.NewBuffer(jsonStr))
	//req, err := http.NewRequest("POST", "/room/create", nil)
	//req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	db, _, _ := createMockUpdateContract()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.UpdateContract)
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
	if out.Message == "Update contracts success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestApiDB_UpdateContractFail2(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
				"id":1,
				"idRoom":1,
				"idOwner":1,
				"idSlave":1,
				"startDate":"abc",
				"endDate":"xyz",
				"circlePay":1,
				"deposit":1,
				"dayPay":1,
				"note":"abc",
				"idBlock":1,
				"userRooms":[
							{
								"idUser":1,
								"idRoom":1}
							]
}`)
	req, err := http.NewRequest("POST", "/room/create", bytes.NewBuffer(jsonStr))
	//req, err := http.NewRequest("POST", "/room/create", nil)
	//req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	//db, _, _ := createMockUpdateContract()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.UpdateContract)
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
	if out.Message == "Update contracts success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
