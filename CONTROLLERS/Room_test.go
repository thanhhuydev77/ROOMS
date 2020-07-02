package CONTROLLERS

import (
	"ROOMS/MODELS"
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createMockGetRoombyId() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	getRoom := sqlmock.NewRows([]string{"Id", "Name", "Floor", "Square", "Price", "Description",
		"IdBlock", "MaxPeople", "Status", "CodeRoom"}).
		AddRow(1, "name1", 1, 1, float64(1), "no descrip", 1, 1, 1, "abc")
	getRoomnil := sqlmock.NewRows([]string{"Id", "Name", "Floor", "Square", "Price", "Description",
		"IdBlock", "MaxPeople", "Status", "CodeRoom"}).AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	mock.ExpectQuery(`SELECT \* FROM ROOMS WHERE id =`).WithArgs(1).WillReturnRows(getRoom)
	mock.ExpectQuery(`SELECT \* FROM ROOMS WHERE id =`).WithArgs(2).WillReturnRows(getRoomnil)
	//mock.ExpectCommit()
	return db, mock, err
}
func TestGetRoomPass(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoombyId()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.GetRoom)
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
	if out.Message != "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomFail(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}
	//req = mux.SetURLVars(req, map[string]string{
	//	"id": "1",
	//})
	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoombyId()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.GetRoom)
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
	if out.Message == "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomFail2(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("GET", "/room/get-rooms", nil)
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
	handler := http.HandlerFunc(a.GetRoom)
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
	if out.Message == "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockGetRoombyblockid() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	getRoom := sqlmock.NewRows([]string{"Id", "Name", "Floor", "Square", "Price", "Description",
		"IdBlock", "MaxPeople", "Status", "CodeRoom"}).
		AddRow(1, "name1", 1, 1, float64(1), "no descrip", 1, 1, 1, "abc")
	mock.ExpectQuery(`SELECT \* FROM ROOMS WHERE idBlock =`).WithArgs(1).WillReturnRows(getRoom)
	//mock.ExpectCommit()
	return db, mock, err
}
func TestGetRoomblockPass(t *testing.T) {
	type data struct {
		Room MODELS.ROOMS `json:"room"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("idBlock", "1")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoombyblockid()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoom)
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
	//fmt.Printf("room :%v",out.Data.Room)
	if out.Message != "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomblockPass2(t *testing.T) {
	type data struct {
		Room MODELS.ROOMS `json:"room"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("idBlock", "1")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoombyblockid()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoom)
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
	//fmt.Printf("room :%v",out.Data.Room)
	if out.Message != "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomblockFail(t *testing.T) {
	type data struct {
		Room MODELS.ROOMS `json:"room"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}

	//q := req.URL.Query()
	//q.Add("idBlock", "1")
	//req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoombyblockid()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoom)
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
	//fmt.Printf("room :%v",out.Data.Room)
	if out.Message == "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomblockFail2(t *testing.T) {
	type data struct {
		Room MODELS.ROOMS `json:"room"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("idBlock", "a")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoombyblockid()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoom)
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
	//fmt.Printf("room :%v",out.Data.Room)
	if out.Message == "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

//
func createMockGetRoombyblockidupdate() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	getRoom := sqlmock.NewRows([]string{"Id", "Nameroom", "Floor", "Square", "Price", "Description",
		"IdBlock", "MaxPeople", "Status", "CodeRoom"}).
		AddRow(1, "name1", 1, 1, float64(1), "no descrip", 1, 1, 1, "abc")
	mock.ExpectQuery(`SELECT R.id, R.nameRoom, R.floor, R.square, R.price, R.description, R.idBlock, R.maxPeople	, R.status, R.codeRoom
FROM ROOMS R LEFT JOIN USER_ROOM UR ON R.id = UR.idRoom WHERE UR.idRoom IS NULL AND idBlock =`).WithArgs(1).WillReturnRows(getRoom)
	//mock.ExpectCommit()
	return db, mock, err
}
func TestGetRoomblockPassupdate(t *testing.T) {
	type data struct {
		Room MODELS.ROOMS `json:"room"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	req, err := http.NewRequest("GET", "/room/get-rooms", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("idBlock", "1")
	q.Add("isMatch", "true")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoombyblockidupdate()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoom)
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
	//fmt.Printf("room :%v",out.Data.Room)
	if out.Message != "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

//
func createMockCreateRoom() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	createRoom := sqlmock.NewRows([]string{"Nameroom", "Floor", "Square", "Price", "Description",
		"IdBlock", "MaxPeople", "Status"}).
		AddRow("name1", 1, 1, float64(1), "no descrip", 1, 1, 1)
	mock.ExpectQuery(`INSERT INTO ROOMS`).WillReturnRows(createRoom)
	//mock.ExpectCommit()
	return db, mock, err
}
func TestCreateRoomPass(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
			"nameRoom": "P101",
			"floor": 1,
			"square": 20,
			"price": 3000000,
			"description": "Phòng 1, 2 phòng ngủ, 5 phòng ăn, tiện nghi đầy đủ chỗ ngủ",
			"idBlock": 51,
			"maxPeople": 10,
			"status": 0
}`)
	req, err := http.NewRequest("POST", "/room/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	//q := req.URL.Query()
	//q.Add("idBlock", "1")
	//q.Add("isMatch","true")
	//req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockCreateRoom()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.CreateRoom)
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
	if out.Message != "Create rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestCreateRoomFail(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
			"nameRoom": "P101",
			"floor": 1,
			"square": 20,
			"price": 3000000,
			"description": "Phòng 1, 2 phòng ngủ, 5 phòng ăn, tiện nghi đầy đủ chỗ ngủ",
			"idBlock": 51,
			"maxPeople": 10,
			"status": 0
}`)
	req, err := http.NewRequest("POST", "/room/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	//q := req.URL.Query()
	//q.Add("idBlock", "1")
	//q.Add("isMatch","true")
	//req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	//db, _, _ := createMockCreateRoom()
	a := &ApiDB{
		nil,
	}

	handler := http.HandlerFunc(a.CreateRoom)
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
	if out.Message == "Create rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestCreateRoomFail2(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{,,,}`)
	req, err := http.NewRequest("POST", "/room/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	//q := req.URL.Query()
	//q.Add("idBlock", "1")
	//q.Add("isMatch","true")
	//req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	//db, _, _ := createMockCreateRoom()
	a := &ApiDB{
		nil,
	}

	handler := http.HandlerFunc(a.CreateRoom)
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
	if out.Message == "Create rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockDeleteRoom() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	DeleteRoom := sqlmock.NewRows([]string{"Nameroom", "Floor", "Square", "Price", "Description",
		"IdBlock", "MaxPeople", "Status"}).
		AddRow("name1", 1, 1, float64(1), "no descrip", 1, 1, 1)
	mock.ExpectQuery(`DELETE FROM ROOMS WHERE id =`).WillReturnRows(DeleteRoom)
	//mock.ExpectCommit()
	return db, mock, err
}
func TestDeleteRoomPass(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/room/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeleteRoom()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.DeleteRoom)
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
	if out.Message != "Delete rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestDeleteRoomFail(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/room/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeleteRoom()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.DeleteRoom)
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
	if out.Message == "Delete rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestDeleteRoomFail2(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/room/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	//db, _, _ := createMockDeleteRoom()
	a := &ApiDB{
		nil,
	}

	handler := http.HandlerFunc(a.DeleteRoom)
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
	if out.Message == "Delete rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockDeletesRoom() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	mock.ExpectExec(`DELETE FROM ROOMS WHERE`).WillReturnResult(sqlmock.NewResult(1, 1))
	//mock.ExpectCommit()
	return db, mock, err
}
func TestDeletesRoomPass(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
       "roomsId": [24, 25]
}`)
	req, err := http.NewRequest("POST", "/bill/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeletesRoom()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.DeleteRooms)
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
	if out.Message != "Delete rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestDeletesRoomFail(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
       "roomsId": [,]
}`)
	req, err := http.NewRequest("POST", "/bill/create", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockDeletesRoom()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.DeleteRooms)
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
	if out.Message == "Delete rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestDeletesRoomFail2(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
       "roomsId": [1,2]
}`)
	req, err := http.NewRequest("POST", "/bill/create", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	//db, _, _ := createMockDeletesRoom()
	a := &ApiDB{
		nil,
	}

	handler := http.HandlerFunc(a.DeleteRooms)
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
	if out.Message == "Delete rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockUpdateRoom() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	DeleteRoom := sqlmock.NewRows([]string{"Nameroom", "Floor", "Square", "Price", "Description",
		"IdBlock", "MaxPeople", "Status"}).
		AddRow("name1", 1, 1, float64(1), "no descrip", 1, 1, 1)
	mock.ExpectQuery(`UPDATE ROOMS`).WillReturnRows(DeleteRoom)
	//mock.ExpectCommit()
	return db, mock, err
}
func TestUpdateRoomPass(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
      	"id": 49,
			"nameRoom": "P101",
			"floor": 1,
			"square": 20,
			"price": 3000000,
			"description": "Phòng 1, 2 phòng ngủ, 5 phòng ăn, tiện nghi đầy đủ chỗ ngủ",
			"idBlock": 51,
			"maxPeople": 10,
			"status": 0,
			"codeRoom": "P1011587213883"
}`)
	req, err := http.NewRequest("POST", "/bill/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	rr := httptest.NewRecorder()
	db, _, _ := createMockUpdateRoom()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.UpdateRoom)
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
	if out.Message != "Update room success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestUpdateRoomFail(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
      	"id": 49,
			"nameRoom": "P101",
			"floor": 1,
			"square": 20,
			"price": 3000000,
			"description": "Phòng 1, 2 phòng ngủ, 5 phòng ăn, tiện nghi đầy đủ chỗ ngủ",
			"idBlock": 51,
			"maxPeople": 10,
			"status": 0,
			"codeRoom": "P1011587213883"
}`)
	req, err := http.NewRequest("POST", "/bill/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	db, _, _ := createMockUpdateRoom()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.UpdateRoom)
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
	if out.Message == "Update room success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestUpdateRoomFail2(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{
      	"id": 49,
			"nameRoom": "P101",
			"floor": 1,
			"square": 20,
			"price": 3000000,
			"description": "Phòng 1, 2 phòng ngủ, 5 phòng ăn, tiện nghi đầy đủ chỗ ngủ",
			"idBlock": 51,
			"maxPeople": 10,
			"status": 0,
			"codeRoom": "P1011587213883"
}`)
	req, err := http.NewRequest("POST", "/bill/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	rr := httptest.NewRecorder()
	//db, _, _ := createMockUpdateRoom()
	a := &ApiDB{
		nil,
	}

	handler := http.HandlerFunc(a.UpdateRoom)
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
	if out.Message == "Update room success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockGetRoomDB() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	DeleteRoom := sqlmock.NewRows([]string{"id", "Nameroom", "Floor", "Square", "Price", "Description",
		"IdBlock", "MaxPeople", "Status", "coderoom", "nameblock", "startdate"}).
		AddRow(1, "name1", 1, 1, float64(1), "no descrip", 1, 1, 1, "dwa", 1, "da")
	mock.ExpectQuery(`SELECT R.*, B.nameBlock, C.startDate FROM`).WillReturnRows(DeleteRoom)
	//mock.ExpectCommit()
	return db, mock, err
}
func TestGetRoomDBPass(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/bill/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("status", "1")
	q.Add("idBlock", "1")
	q.Add("userId", "1")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoomDB()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoomDB)
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
	if out.Message != "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomDBFail(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/bill/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	//q.Add("status", "1")
	q.Add("idBlock", "1")
	q.Add("userId", "1")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoomDB()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoomDB)
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
	if out.Message == "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomDBFail3(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/bill/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("status", "1")
	q.Add("idBlock", "1")
	//q.Add("userId", "1")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoomDB()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoomDB)
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
	if out.Message == "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomDBFail2(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/bill/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("status", "1")
	//q.Add("idBlock", "1")
	q.Add("userId", "1")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoomDB()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoomDB)
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
	if out.Message == "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomDBFail4(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/bill/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("status", "a")
	q.Add("idBlock", "1")
	q.Add("userId", "1")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoomDB()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoomDB)
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
	if out.Message == "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomDBFail5(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/bill/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("status", "1")
	q.Add("idBlock", "1")
	q.Add("userId", "1")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	//db, _, _ := createMockGetRoomDB()
	a := &ApiDB{
		nil,
	}

	handler := http.HandlerFunc(a.GetRoomDB)
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
	if out.Message == "Get rooms success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockGetRoomImage() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	DeleteRoom := sqlmock.NewRows([]string{"Id", "Name", "Status", "Url", "CodeRoom"}).
		AddRow(1, "av", "va", "av", "av")
	mock.ExpectQuery(`SELECT \* FROM ROOM_IMAGES`).WillReturnRows(DeleteRoom)
	//mock.ExpectCommit()
	return db, mock, err
}
func TestGetRoomImagePass(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/bill/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("codeRoom", "1")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoomImage()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoomImage)
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
	if out.Message != "Get images success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomImageFail(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/bill/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	//q := req.URL.Query()
	//q.Add("codeRoom", "1")
	//req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoomImage()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoomImage)
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
	if out.Message == "Get images success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomImageFail2(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/bill/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("codeRoom", "1")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	//db, _, _ := createMockGetRoomImage()
	a := &ApiDB{
		nil,
	}

	handler := http.HandlerFunc(a.GetRoomImage)
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
	if out.Message == "Get images success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockGetRoomUser() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	DeleteRoom := sqlmock.NewRows([]string{"RoomName"}).
		AddRow("abc")
	mock.ExpectQuery(`SELECT C.fullName FROM USER_ROOM`).WillReturnRows(DeleteRoom)
	//mock.ExpectCommit()
	return db, mock, err
}
func TestGetRoomUserPass(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/bill/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("idRoom", "1")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoomUser()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoomUser)
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
	if out.Message != "Get user rent success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomUserFail(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/bill/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	//q := req.URL.Query()
	//q.Add("idRoom", "1")
	//req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoomUser()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoomUser)
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
	if out.Message == "Get user rent success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomUserFail3(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/bill/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("idRoom", "ac")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	db, _, _ := createMockGetRoomUser()
	a := &ApiDB{
		db,
	}

	handler := http.HandlerFunc(a.GetRoomUser)
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
	if out.Message == "Get user rent success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestGetRoomUserFail2(t *testing.T) {

	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	req, err := http.NewRequest("POST", "/bill/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("idRoom", "1")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	//db, _, _ := createMockGetRoomUser()
	a := &ApiDB{
		nil,
	}

	handler := http.HandlerFunc(a.GetRoomUser)
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
	if out.Message == "Get user rent success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
