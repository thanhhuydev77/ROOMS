package CONTROLLERS

import (
	"ROOMS/BUSINESS"
	"ROOMS/MODELS"
	. "ROOMS/STATICS"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"time"
)

func TokenHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	r.ParseForm()

	UserName := r.Form.Get("userName")
	Pass := r.Form.Get("pass")

	IsExsist, passok, a := BUSINESS.Login(UserName, Pass)

	if !IsExsist {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"message": "Can't find user please sign in again!"}`)
		return
	}

	if !passok {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"message": "Your password is wrong, please type again !"}`)
		return
	}
	// expired after 15 dates
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": a.UserName,
		"exp":  time.Now().Add(time.Hour * time.Duration(15*24)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(APP_KEY))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"token_generation_failed"}`)
		return
	}
	//w.Write([]byte(`{"hello": "world"}`))
	stringresult := `{"message": "Login success","data":{"token":"` + tokenString + `","user":{ "id":` + strconv.Itoa(a.Id) + `,
						"username":"` + a.UserName + `","fullname":"` + a.FullName.String + `","role":` + strconv.Itoa(int(a.Role.Int32)) + `}}}`
	io.WriteString(w, stringresult)
	return
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	User := MODELS.USERS{}
	r.ParseForm()
	// parse user information
	User.UserName = r.Form.Get("userName")
	User.Pass = r.Form.Get("pass")
	User.FullName.String = r.Form.Get("fullName")
	//User.IdentifyFront = r.Form.Get("identifyFront")
	//User.IdentifyBack = r.Form.Get("identifyBack")
	//DateBirth, err := time.Parse("01/02/2006", r.Form.Get("dateBirth"))
	//if err != nil {
	//	w.WriteHeader(http.StatusBadGateway)
	//	io.WriteString(w, `{"message": "can not parse datebirth!"}`+err.Error())
	//	return
	//}
	//User.DateBirth = DateBirth

	User.Address.String = r.Form.Get("address")
	Role, err := strconv.Atoi(r.Form.Get("role"))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, `{"message": "can not parse role!"}`)
		return
	}
	User.Role.Int32 = int32(Role)
	User.Sex.String = r.Form.Get("sex")
	//User.Job = r.Form.Get("job")
	//TempReg, err := strconv.Atoi(r.Form.Get("tempReg"))
	//if err != nil {
	//	w.WriteHeader(http.StatusBadGateway)
	//	io.WriteString(w, `{"message": "can not parse Tempreg!"}`)
	//	return
	//}
	//User.TempReg = TempReg
	User.Province.String = r.Form.Get("province")
	User.Email.String = r.Form.Get("email")
	confirm := r.Form.Get("confirm")
	if confirm != User.Pass {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message": "pass and confirm must be same!"}`)
		return
	}

	result, err := BUSINESS.Register(User)
	if result {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"message": "Register success","data": {"status": 1}}`)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":{"code":"`+err.Error()+`"}}`)
	}
}

func GetallUserName(w http.ResponseWriter, r *http.Request) {

	// Query()["key"] will return an array of items,
	// we only want the single item.

	allusername := BUSINESS.GetAllUserName()
	w.Header().Add("Content-Type", "application/json")
	if allusername == nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message": "unsuccess"}`)
		return
	}
	w.WriteHeader(200)
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
	stringresult := `{"message": "Get Users success","status": 200,"data":{`
	for _, val := range List {
		stringresult += val.String() + ","
	}
	stringresult += "}}"
	io.WriteString(w, stringresult)
}
