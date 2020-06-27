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

func createMockUnit() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	getallunit := sqlmock.NewRows([]string{"id", "name", "description"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")
	mock.ExpectQuery(`select \* from UNITS`).WillReturnRows(getallunit)

	//mock.ExpectCommit()
	return db, mock, err
}

func TestGetAllUnitPass(t *testing.T) {

	type data struct {
		Units []MODELS.UNIT `json:"units"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/unit/get-units", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	db, _, _ := createMockUnit()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.GetAllUnit)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	out := output{
		Data: data{nil},
	}
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

	type data struct {
		Units []MODELS.UNIT `json:"units"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/unit/get-units", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	//db, _, _ := createMockUnit()

	a := &ApiDB{
		nil,
	}

	handler := http.HandlerFunc(a.GetAllUnit)
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
	if out.Message == "Get units success" || len(out.Data.Units) > 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Units))
	}
}
