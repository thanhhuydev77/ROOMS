package CONTROLLERS

import (
	"ROOMS/MODELS"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
)

func createMockBlock() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	result := sqlmock.NewRows([]string{"id", "nameBlock", "address", "description", "idowner"}).
		AddRow(68, "A3", "KTX Khu A - ĐHQG HCM", "To nhất quả đất", 69).
		AddRow(69, "A2", "KTX Khu A", "To", 69)
	mock.ExpectQuery(`select \* from BLOCKS where idowner = .*`).WillReturnRows(result)

	return db, mock, err
}

func TestGetBlocksPass(t *testing.T) {

	type data struct {
		Blocks []MODELS.BLOCKS `json:"blocks"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/block/get-block", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("userId", "69")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockBlock()
	a := &ApiDB{
		db,
	}

	handle := http.HandlerFunc(a.GetBlock)
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

	if out.Message != "Get Blocks success" || len(out.Data.Blocks) == 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Blocks))
	}
}

func TestGetBlocksFail1(t *testing.T) {

	type data struct {
		Blocks []MODELS.BLOCKS `json:"blocks"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/block/get-block", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("userId", "68")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	// db, _, _ := createMockBlock()
	a := &ApiDB{
		nil,
	}

	handle := http.HandlerFunc(a.GetBlock)
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

	if out.Message != "Get Blocks success" || len(out.Data.Blocks) > 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Blocks))
	}
}

func TestGetBlocksFail2(t *testing.T) {

	type data struct {
		Blocks []MODELS.BLOCKS `json:"blocks"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/block/get-block", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("userId", "")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	a := &ApiDB{
		nil,
	}

	handle := http.HandlerFunc(a.GetBlock)
	handle.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func createMockBlockID() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	resultID := sqlmock.NewRows([]string{"id", "nameBlock", "address", "description", "idowner"}).
		AddRow(68, "A3", "KTX Khu A - ĐHQG HCM", "To nhất quả đất", 69)
	mock.ExpectQuery(`select \* from BLOCKS where id = \?`).WillReturnRows(resultID)

	return db, mock, err
}
func TestGetBlockByIdPass(t *testing.T) {

	type data struct {
		Blocks MODELS.BLOCKS `json:"blocks"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/block/get-block/68", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "68",
	})

	rr := httptest.NewRecorder()
	db, _, _ := createMockBlockID()
	a := &ApiDB{
		db,
	}

	handle := http.HandlerFunc(a.GetBlockById)
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

	if out.Message != "Get Blocks success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func TestGetBlockIdFail1(t *testing.T) {

	type data struct {
		Blocks MODELS.BLOCKS `json:"blocks"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/block/get-block/60", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "60",
	})

	rr := httptest.NewRecorder()
	a := &ApiDB{
		nil,
	}

	handle := http.HandlerFunc(a.GetBlockById)
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

	if out.Message != "Get Blocks success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func TestGetBlockIdFail2(t *testing.T) {

	type data struct {
		Blocks MODELS.BLOCKS `json:"blocks"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/block/get-block/60", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"idd": "60",
	})

	rr := httptest.NewRecorder()
	a := &ApiDB{
		nil,
	}

	handle := http.HandlerFunc(a.GetBlockById)
	handle.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
