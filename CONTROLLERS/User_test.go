package CONTROLLERS

import (
	"ROOMS/MODELS"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createMockUserLogin() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	login := sqlmock.NewRows([]string{"id", "userName", "pass", "fullName", "role"}).
		AddRow(1, "test", "$2a$05$p8ae8Ugx/IEDhMa6s6PTTebEmD/6GxCQ2oVOpKJvxz18tzwgIKPIO", "fullname1", 1)
	mock.ExpectQuery(`select id,username,pass,fullname,role from USERS where username = \?`).WillReturnRows(login)
	return db, mock, err
}
func TestLoginPass(t *testing.T) {

	type data struct {
		Token string       `json:"token"`
		User  MODELS.USERS `json:"user"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}
	var jsonStr = []byte(`{"userName":"test", "pass":"hash"}`)
	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	db, _, _ := createMockUserLogin()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.UserLogin)
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

	if out.Message != "Login success" || len(out.Data.Token) == 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Token))
	}
}
func TestLoginFail(t *testing.T) {

	type data struct {
		Token string       `json:"token"`
		User  MODELS.USERS `json:"user"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	var jsonStr = []byte(`{"userName":"test", "pass":"hash"}`)
	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	//db, _, _ := createMockDS()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.UserLogin)
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
	if out.Message == "Login success" || len(out.Data.Token) > 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Token))
	}
}
func TestLoginFail3(t *testing.T) {

	type data struct {
		Token string       `json:"token"`
		User  MODELS.USERS `json:"user"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	var jsonStr = []byte(`{""}`)
	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	db, _, _ := createMockUserLogin()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.UserLogin)
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
	if out.Message == "Login success" || len(out.Data.Token) > 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Token))
	}
}
func TestLoginFail2(t *testing.T) {

	type data struct {
		Token string       `json:"token"`
		User  MODELS.USERS `json:"user"`
	}
	type output struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	var jsonStr = []byte(`{"userName":"test", "pass":"hash1"}`)
	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	db, _, _ := createMockUserLogin()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.UserLogin)
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
	if out.Message == "Login success" || len(out.Data.Token) > 0 {
		t.Errorf("error message(%v) or lendata(%v)", out.Message, len(out.Data.Token))
	}
}

func createMockUserRegister() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	Register := sqlmock.NewRows([]string{"id", "userName", "pass", "fullName", "role"}).
		AddRow(1, "test", "$2a$05$p8ae8Ugx/IEDhMa6s6PTTebEmD/6GxCQ2oVOpKJvxz18tzwgIKPIO", "fullname1", 1)
	mock.ExpectQuery("insert into USERS.*").WillReturnRows(Register)
	return db, mock, err
}
func TestRegisterPass(t *testing.T) {

	type output struct {
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{	"userName":"huydz123","pass":"hash","confirm":"hash","fullName":"fullname","sex":"nam",
	"email":"abc@gm.ddd","role":1,"province":"binhdinh","address":"abcdad"}`)
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	db, _, _ := createMockUserRegister()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.UserRegister)
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

	if out.Message != "Register success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestRegisterFail(t *testing.T) {

	type output struct {
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{	"userName":"huydz123","pass":"hash","confirm":"hash","fullName":"fullname","sex":"nam",
	"email":"abc@gm.ddd","role":1,"province":"binhdinh","address":"abcdad"}`)
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.UserRegister)
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

	if out.Message == "Register success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestRegisterFail2(t *testing.T) {

	type output struct {
		Message string `json:"message"`
	}
	var jsonStr = []byte(``)
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.UserRegister)
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

	if out.Message == "Register success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockUserName() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	username := sqlmock.NewRows([]string{"userName"}).
		AddRow("huydz da test")
	mock.ExpectQuery("select username from USERS").WillReturnRows(username)
	return db, mock, err
}
func TestApiDB_GetallUserNamePass(t *testing.T) {

	type output struct {
		Message string   `json:"message"`
		Data    []string `json:"data"`
	}

	req, err := http.NewRequest("POST", "/user/get-all-username", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	db, _, _ := createMockUserName()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.GetallUserName)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var out output
	err = json.Unmarshal(rr.Body.Bytes(), &out)
	if err != nil {
		s := string(rr.Body.Bytes())
		fmt.Println(s) // ABC€
		t.Errorf("error marshal :%v", err)
	}

	if out.Message != "get all username success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestApiDB_GetallUserNameFail(t *testing.T) {

	type output struct {
		Message string   `json:"message"`
		Data    []string `json:"data"`
	}

	req, err := http.NewRequest("POST", "/user/get-all-username", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	//db, _, _ := createMockUserName()
	a := &ApiDB{
		nil,
	}
	handler := http.HandlerFunc(a.GetallUserName)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var out output
	err = json.Unmarshal(rr.Body.Bytes(), &out)
	if err != nil {
		s := string(rr.Body.Bytes())
		fmt.Println(s) // ABC€
		t.Errorf("error marshal :%v", err)
	}

	if out.Message == "get all username success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

func createMockGetUser() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	user := sqlmock.NewRows([]string{"Id", "UserName", "Pass", "FullName", "IdentifyFro", "IdentifyBac", "DateBirth",
		"Address", "Role", "Sex", "Job", "WorkPlace", "TempReg", "Province", "Email", "Avatar", "PhoneNumber"}).
		AddRow(1, "HUYDZ", "abc", "huydzzz", 1, 1, nil, "say hi", 1, "nam", "ranh", "cau thi nghe", 0, "dalat", "abcmal", "dw", "dwad")
	mock.ExpectQuery(`select \* from USERS`).WillReturnRows(user)
	return db, mock, err
}
func TestApiDB_GetUser1(t *testing.T) {

	type output struct {
		Message string       `json:"message"`
		Status  int          `json:"status"`
		Data    MODELS.USERS `json:"data"`
	}

	req, err := http.NewRequest("POST", "/user/get-user", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	db, _, _ := createMockGetUser()
	a := &ApiDB{
		db,
	}
	handler := http.HandlerFunc(a.GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var out output
	err = json.Unmarshal(rr.Body.Bytes(), &out)
	if err != nil {
		s := string(rr.Body.Bytes())
		fmt.Println(s) // ABC€
		t.Errorf("error marshal :%v", err)
	}

	if out.Message != "Get Users success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
