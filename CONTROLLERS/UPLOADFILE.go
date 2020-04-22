package CONTROLLERS

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadPicture(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./ASSET/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	io.Copy(f, file)
}
