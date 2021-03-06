package CONTROLLERS

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

//upload a image
func UploadPicture(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	file, handler, err := r.FormFile("file")
	if err != nil {
		//fmt.Print(err)
		io.WriteString(w, `{"message":"Can’t upload avatar"}`)
		return
	}
	defer file.Close()

	f, _ := os.OpenFile("../public/images/avatars/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	//if err != nil {
	//	//fmt.Printf("error :%v",err)
	//	io.WriteString(w, `{"message":"Can’t upload avatar"}`)
	//	return
	//}
	defer f.Close()
	io.Copy(f, file)
	io.WriteString(w, `{ "status": 200,
    "message": "Upload avatar success",
    "data": {
        "fieldname": "file",
        "originalname": "`+handler.Filename+`",
        "destination": "public",
		 "mimetype": "`+handler.Header.Get("Content-Type")+`",
        "filename": "`+handler.Filename+`",
        "path": "public\\images\\avatars\\`+handler.Filename+`",
        "size": `+strconv.Itoa(int(handler.Size))+`
    }
}`)
}
