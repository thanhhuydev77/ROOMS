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

func createMockDS() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	getallDS := sqlmock.NewRows([]string{"id", "nameService", "price", "description", "idUnit", "unitName"}).
		AddRow(1, "service1", 1, "hello", 1, "unitname1")
	mock.ExpectQuery(`SELECT DS.*,U.name FROM DEFAULT_SERVICES as DS INNER JOIN UNITS as U on DS.idUnit = U.id`).WillReturnRows(getallDS)

	return db, mock, err
}

func TestGetAllDSPass(t *testing.T) {

	type data struct {
		DefaultServices []MODELS.GET_DEFAULT_SERVICES_REQUEST `json:"defaultServices"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/default-service/get-default-services", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	db, _, _ := createMockDS()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.Get_default_service)
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
	if out.Message != "Get default services success" || len(out.Data.DefaultServices) == 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.DefaultServices))
	}
}

func TestGetAllDSFail(t *testing.T) {

	type data struct {
		DefaultServices []MODELS.GET_DEFAULT_SERVICES_REQUEST `json:"defaultServices"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/default-service/get-default-services", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	//db, _, _ := createMockDS()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.Get_default_service)
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
	if out.Message == "Get default services success" || len(out.Data.DefaultServices) > 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.DefaultServices))
	}
}
