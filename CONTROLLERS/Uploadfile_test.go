package CONTROLLERS

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, _ := http.NewRequest("POST", uri, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	return request, nil
}
func TestApiDB_UploadFilePass(t *testing.T) {
	//type data struct {
	//	Fieldname string	`json:""`
	//	Originalname string	`json:""`
	//	Destination string	`json:""`
	//	Mimetype string		`json:""`
	//	Filename string		`json:""`
	//	Path string			`json:""`
	//	Size string			`json:""`
	//}
	type output struct {
		Message string `json:"message"`
		//Data    []string `json:"data"`
	}
	extraParams := map[string]string{
		"title":       "My Document",
		"author":      "Matt Aimonetti",
		"description": "A document with all the Go programming language secrets",
	}
	req, err := newfileUploadRequest("/upload/userAvatar", extraParams, "file", `C:\Users\Truong Duy\Desktop\duy.jpg`)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(UploadPicture)
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

	if out.Message != "Upload avatar success" {
		t.Errorf("error message(%v)", out.Message)
	}
}
func TestApiDB_UploadFileFail1(t *testing.T) {
	//type data struct {
	//	Fieldname string	`json:""`
	//	Originalname string	`json:""`
	//	Destination string	`json:""`
	//	Mimetype string		`json:""`
	//	Filename string		`json:""`
	//	Path string			`json:""`
	//	Size string			`json:""`
	//}
	type output struct {
		Message string `json:"message"`
		//Data    []string `json:"data"`
	}

	req, err := http.NewRequest("POST", "/upload/userAvatar", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(UploadPicture)
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

	if out.Message == "Upload avatar success" {
		t.Errorf("error message(%v)", out.Message)
	}
}

//func TestApiDB_UploadFileFail2(t *testing.T) {
//	//type data struct {
//	//	Fieldname string	`json:""`
//	//	Originalname string	`json:""`
//	//	Destination string	`json:""`
//	//	Mimetype string		`json:""`
//	//	Filename string		`json:""`
//	//	Path string			`json:""`
//	//	Size string			`json:""`
//	//}
//	type output struct {
//		Message string   `json:"message"`
//		//Data    []string `json:"data"`
//	}
//	extraParams := map[string]string{
//		"title":       "My Document",
//		"author":      "Matt Aimonetti",
//		"description": "A document with all the Go programming language secrets",
//	}
//	req, err := newfileUploadRequest("/upload/userAvatar",extraParams, "file", `F:\download\abc.jpg`)
//	if err != nil {
//		t.Fatal(err)
//	}
//	rr := httptest.NewRecorder()
//
//	handler := http.HandlerFunc(UploadPicture)
//	handler.ServeHTTP(rr, req)
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//
//	var out output
//	err = json.Unmarshal(rr.Body.Bytes(), &out)
//	if err != nil {
//		//s := string(rr.Body.Bytes())
//		//fmt.Println(s) // ABC€
//		t.Errorf("error marshal :%v", err)
//	}
//
//	if out.Message == "Upload avatar success" {
//		t.Errorf("error message(%v)", out.Message)
//	}
//}
