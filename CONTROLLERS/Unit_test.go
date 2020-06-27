package CONTROLLERS

import (
	"ROOMS/MODELS"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"
	"testing"
)

var db *sql.DB = nil
var mock sqlmock.Sqlmock = nil
var err error = nil

func GetMockDb() (*sql.DB, sqlmock.Sqlmock, error) {
	if db == nil || mock == nil {
		db, mock, err = createMockDb()
	}
	return db, mock, err
}
func createMockDb() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	getallunit := sqlmock.NewRows([]string{"id", "name", "description"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")

	mock.ExpectQuery(`select \* from UNITS`).WillReturnRows(getallunit)

	return db, mock, err
}

type data struct {
	Units []MODELS.UNIT `json:"units"`
}
type Output struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    data   `json:"data"`
}

func TestGetAllUnitPass(t *testing.T) {
	req, err := http.NewRequest("GET", "/unit/get-units", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	db, _, _ = GetMockDb()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.GetAllUnit)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var out Output
	err = json.Unmarshal(rr.Body.Bytes(), &out)
	if err != nil {
		t.Errorf("error marshal :%v", err)
	}
	if out.Message != "Get units success" || len(out.Data.Units) == 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Units))
	}
}

//
func TestGetAllUnitFail(t *testing.T) {
	req, err := http.NewRequest("GET", "/unit/get-units", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	db, _, _ = GetMockDb()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.GetAllUnit)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var out Output
	err = json.Unmarshal(rr.Body.Bytes(), &out)
	if err != nil {
		t.Errorf("error marshal :%v", err)
	}
	if out.Message == "Get units success" || len(out.Data.Units) > 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Units))
	}
}
