package CONTROLLERS

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthTokenFail(t *testing.T) {

	type output struct {
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{	"token":""}`)
	req, err := http.NewRequest("POST", "/user/validate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := AuthMiddleware(http.HandlerFunc(ValidateToken))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var out output
	err = json.Unmarshal(rr.Body.Bytes(), &out)
	if err == nil {
		t.Errorf("error marshal :%v", err)
	}
	if out.Message == "Validate success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestAuthTokenPass(t *testing.T) {

	type output struct {
		Message string `json:"message"`
	}
	var jsonStr = []byte(`{	"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU3NDA0MTIsImlhdCI6MTU4OTM0MDQxMiwidXNlciI6InRlc3QifQ.qC_MOzb41P4qO1QMNFNXYHLFnT8TDMEMURurr7eFH1c"}`)
	req, err := http.NewRequest("POST", "/user/validate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ValidateToken)
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
	if out.Message != "Validate success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
