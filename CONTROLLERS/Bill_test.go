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

func createMockGetBill() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	getBill := sqlmock.NewRows([]string{"id", "IdRoom", "DateCheckOut", "TotalPrice", "IsCheckedOut"}).
		AddRow(1, 1, nil, 1, 1)
	mock.ExpectQuery(`SELECT \* FROM BILLS .*`).WillReturnRows(getBill)

	getBilldt := sqlmock.NewRows([]string{"id", "IdBill", "IdService", "Amount", "TotalPrice"}).
		AddRow(1, 1, 1, 1, 1)
	mock.ExpectQuery(`SELECT \* FROM BILL_DETAILS where idBill = .*`).WillReturnRows(getBilldt)
	return db, mock, err
}
func TestGetBillPass(t *testing.T) {

	type Data struct {
		Bill       MODELS.BILLS          `json:"bill"`
		Billdetail []MODELS.BILL_DETAILS `json:"billdetail"`
	}
	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/bill/get-bill-by-id/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	rr := httptest.NewRecorder()
	db, _, _ := createMockGetBill()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.GetBills)
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
	if out.Message != "get bill success" || len(out.Data.Billdetail) == 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Billdetail))
	}
}
func TestGetBillFail(t *testing.T) {

	type Data struct {
		Bill       MODELS.BILLS          `json:"bill"`
		Billdetail []MODELS.BILL_DETAILS `json:"billdetail"`
	}
	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/bill/get-bill-by-id/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	//req = mux.SetURLVars(req, map[string]string{
	//	"id": "1",
	//})
	rr := httptest.NewRecorder()
	db, _, _ := createMockGetBill()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.GetBills)
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
	if out.Message == "get bill success" || len(out.Data.Billdetail) > 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Billdetail))
	}
}
func TestGetBillFail2(t *testing.T) {

	type Data struct {
		Bill       MODELS.BILLS          `json:"bill"`
		Billdetail []MODELS.BILL_DETAILS `json:"billdetail"`
	}
	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/bill/get-bill-by-id/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	rr := httptest.NewRecorder()
	//db, _, _ := createMockGetBill()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.GetBills)
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
	if out.Message == "get bill success" || len(out.Data.Billdetail) > 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Billdetail))
	}
}

func TestGetBillBlockPass(t *testing.T) {

	type Data struct {
		Bills []MODELS.BILLS `json:"bills"`
	}

	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/bill/get-bills", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	rr := httptest.NewRecorder()
	db, _, _ := createMockGetBill()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.GetBillsbyblock)
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
	if out.Message != "get bill success" || len(out.Data.Bills) == 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Bills))
	}
}
func TestGetBillBlockFail(t *testing.T) {
	type Data struct {
		Bills []MODELS.BILLS `json:"bills"`
	}

	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/bill/get-bill-by-id/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	//req = mux.SetURLVars(req, map[string]string{
	//	"id": "1",
	//})
	rr := httptest.NewRecorder()
	db, _, _ := createMockGetBill()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.GetBillsbyblock)
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
	if out.Message == "get bill success" || len(out.Data.Bills) > 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Bills))
	}
}
func TestGetBillBlockFail2(t *testing.T) {

	type Data struct {
		Bills []MODELS.BILLS `json:"bills"`
	}

	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/bill/get-bill-by-id/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	rr := httptest.NewRecorder()
	//db, _, _ := createMockGetBill()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.GetBillsbyblock)
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
	if out.Message == "get bill success" || len(out.Data.Bills) > 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Bills))
	}
}

func createMockCreateBill() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	//abc := interface{}
	insertBill := sqlmock.NewRows([]string{"id"}).
		AddRow(1)
	getBillid := sqlmock.NewRows([]string{"id"}).
		AddRow(1)
	mock.ExpectBegin()
	mock.ExpectQuery(`insert into BILLS.*`).WithArgs(49, "1", float64(1), 1).WillReturnRows(insertBill)
	mock.ExpectQuery(`select id from BILLS where .*`).WillReturnRows(getBillid)
	mock.ExpectQuery("INSERT INTO BILL_DETAILS.*").WithArgs(1, 78, 1, float64(2)).WillReturnRows(insertBill)
	mock.ExpectCommit()
	return db, mock, err
}
func TestCreateBillPass(t *testing.T) {

	type Data struct {
		Status int `json:"status"`
	}
	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}
	var jsonStr = []byte(`{
        "idRoom":49,
        "dateCheckOut": "1",
        "totalPrice": 1,
        "isCheckedOut": 1,
        "billDetail": [
            {
                "idService": 78,
                "amount": 1,
                "totalPrice": 2
            }
        ]
}`)
	req, err := http.NewRequest("POST", "/bill/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockCreateBill()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.CreateBill)
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
	if out.Message != "Create bill success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestCreateBillFail(t *testing.T) {

	type Data struct {
		Status int `json:"status"`
	}
	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}
	var jsonStr = []byte(`{
        "idRoom":49,
        "dateCheckOut": "1",
        "totalPrice": 1,
        "isCheckedOut": 1,
        "billDetail": [
            {
                "idService": 78,
                "amount": 1,
                "totalPrice": 2
            }
        ]
}`)
	req, err := http.NewRequest("POST", "/bill/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	//db, _, _ := createMockCreateBill()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.CreateBill)
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
	if out.Message == "Create bill success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestCreateBillFail2(t *testing.T) {

	type Data struct {
		Status int `json:"status"`
	}
	type Result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}
	var jsonStr = []byte(`{,,}`)
	req, err := http.NewRequest("POST", "/bill/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockCreateBill()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.CreateBill)
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
	if out.Message == "Create bill success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockDeleteBill() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	mock.ExpectExec("DELETE FROM BILLS.*").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	return db, mock, err
}
func TestDeleteBillPass(t *testing.T) {

	type Result struct {
		Message string `json:"message"`
	}

	req, err := http.NewRequest("DELETE", "/bill/delete", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeleteBill()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.DeleteBill)
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
	if out.Message != "Delete bill success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestDeleteBillFail(t *testing.T) {

	type Result struct {
		Message string `json:"message"`
	}

	req, err := http.NewRequest("DELETE", "/bill/delete", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	//db, _, _ := createMockDeleteBill()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.DeleteBill)
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
	if out.Message == "Create bill success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestDeleteBillFail2(t *testing.T) {

	type Result struct {
		Message string `json:"message"`
	}

	req, err := http.NewRequest("DELETE", "/bill/delete", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	//db, _, _ := createMockDeleteBill()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.DeleteBill)
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
	if out.Message == "Create bill success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockUpdateBill() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE BILLS SET.*`).WithArgs("1", float64(1), 1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`DELETE FROM BILL_DETAILS.*`).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO BILL_DETAILS.*").WithArgs(1, 78, 1, float64(2)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	return db, mock, err
}
func TestUpdateBillPass(t *testing.T) {

	type Result struct {
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
        "idRoom":49,
        "dateCheckOut": "1",
        "totalPrice": 1,
        "isCheckedOut": 1,
        "billDetail": [
            {
                "idService": 78,
                "amount": 1,
                "totalPrice": 2
            }
        ]
}`)
	req, err := http.NewRequest("DELETE", "/bill/update/1", bytes.NewBuffer(jsonStr))
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockUpdateBill()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.UpdateBill)
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
	if out.Message != "Update bill success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestUpdateBillFail(t *testing.T) {

	type Result struct {
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{,,}`)
	req, err := http.NewRequest("DELETE", "/bill/update/1", bytes.NewBuffer(jsonStr))
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockUpdateBill()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.UpdateBill)
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
	if out.Message == "Update bill success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestUpdateBillFail1(t *testing.T) {

	type Result struct {
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
        "idRoom":49,
        "dateCheckOut": "1",
        "totalPrice": 1,
        "isCheckedOut": 1,
        "billDetail": [
            {
                "idService": 78,
                "amount": 1,
                "totalPrice": 2
            }
        ]
}`)
	req, err := http.NewRequest("DELETE", "/bill/update/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockUpdateBill()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.UpdateBill)
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
	if out.Message == "Update bill success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestUpdateBillFail2(t *testing.T) {

	type Result struct {
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
        "idRoom":49,
        "dateCheckOut": "1",
        "totalPrice": 1,
        "isCheckedOut": 1,
        "billDetail": [
            {
                "idService": 78,
                "amount": 1,
                "totalPrice": 2
            }
        ]
}`)
	req, err := http.NewRequest("DELETE", "/bill/update/1", bytes.NewBuffer(jsonStr))
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	//db, _, _ := createMockUpdateBill()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.UpdateBill)
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
	if out.Message == "Update bill success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
