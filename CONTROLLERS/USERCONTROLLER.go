package CONTROLLERS

import (
	"ROOMS/BUSINESS"
	"ROOMS/MODELS"
	. "ROOMS/STATICS"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"time"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	p := MODELS.USERS{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		io.WriteString(w, `{"message": "wrong format!"}`)
		return
	}

	IsExsist, passok, a := BUSINESS.Login(p.UserName, p.Pass)

	if !IsExsist {
		//w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"message": "Can't find user please sign in again!"}`)
		return
	}

	if !passok {
		//w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"message": "Your password is wrong, please type again !"}`)
		return
	}
	// expired after 1000 dates
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": a.UserName,
		"exp":  time.Now().Add(time.Hour * time.Duration(1000*24)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(APP_KEY))
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"token_generation_failed"}`)
		return
	}

	//w.Write([]byte(`{"hello": "world"}`))
	stringresult := `{"message": "Login success","data":{"token":"` + tokenString + `","user":{ "id":` + strconv.Itoa(a.Id) + `,
						"username":"` + a.UserName + `","fullname":"` + a.FullName.String + `","role":` + strconv.Itoa(int(a.Role.Int64)) + `}}}`
	io.WriteString(w, stringresult)
	return
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	p := MODELS.RequestRegister{}
	err1 := json.NewDecoder(r.Body).Decode(&p)
	if err1 != nil {
		io.WriteString(w, `{"message": "wrong format!"}`+err1.Error())
		return
	}

	result, err := BUSINESS.Register(p)
	if result {
		io.WriteString(w, `{"message": "Register success","data": {"status": 1}}`)
	} else {
		io.WriteString(w, `{"message":{"code":"`+err.Error()+`"}}`)
	}
}

func GetallUserName(w http.ResponseWriter, r *http.Request) {
	// Query()["key"] will return an array of items,
	// we only want the single item.

	allusername := BUSINESS.GetAllUserName()
	w.Header().Add("Content-Type", "application/json")
	if allusername == nil {
		//w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message": "unsuccess"}`)
		return
	}
	//w.WriteHeader(200)
	io.WriteString(w, `{"message": "success","data":{`)
	for _, val := range allusername {
		io.WriteString(w, "\""+val+"\",")
	}
	io.WriteString(w, "}}")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	Id, err := strconv.Atoi(vars["Id"])
	//have not id --> get all
	if err != nil {
		Id = -1
	}
	List := BUSINESS.GetUsers(Id)
	jsonlist, err1 := json.Marshal(List)
	result := List[0]
	if err1 != nil {
		return
	}
	stringresult := `{"message": "Get Users success","status": 200,"data":{"user":`
	if len(List) == 1 {
		jsonresult, _ := json.Marshal(result)
		stringresult += string(jsonresult)
	} else {
		stringresult += string(jsonlist)
	}
	stringresult += "}}"

	io.WriteString(w, stringresult)
}
