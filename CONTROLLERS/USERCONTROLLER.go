package CONTROLLERS

import (
	"ROOMS/BUSINESS"
	"ROOMS/MODELS"
	. "ROOMS/STATICS"
	"github.com/dgrijalva/jwt-go"
	"io"
	"net/http"
	"strconv"
	"time"
)

func TokenHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
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
						"username":"` + a.UserName + `","fullname":"` + a.FullName + `","role":` + strconv.Itoa(a.Role) + `}}}`
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
	User.FullName = r.Form.Get("fullName")
	User.IdentifyFront = r.Form.Get("identifyFront")
	User.IdentifyBack = r.Form.Get("identifyBack")
	DateBirth, err := time.Parse("01/02/2006", r.Form.Get("dateBirth"))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, `{"message": "can not parse datebirth!"}`+err.Error())
		return
	}
	User.DateBirth = DateBirth

	User.Address = r.Form.Get("address")
	Role, err := strconv.Atoi(r.Form.Get("role"))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, `{"message": "can not parse role!"}`)
		return
	}
	User.Role = Role
	User.Sex = r.Form.Get("sex")
	User.Job = r.Form.Get("job")
	TempReg, err := strconv.Atoi(r.Form.Get("tempReg"))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, `{"message": "can not parse Tempreg!"}`)
		return
	}
	User.TempReg = TempReg
	User.Province = r.Form.Get("province")
	User.Email = r.Form.Get("email")

	if BUSINESS.Register(User) {
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, `{"message": "Register success"}`)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message": "Register unsuccess"}`)
	}

}

func GetallUserName(w http.ResponseWriter, r *http.Request) {

	allusername := BUSINESS.GetAllUserName()
	w.Header().Add("Content-Type", "application/json")
	if allusername == nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message": "unsuccess"}`)
		return
	}
	w.WriteHeader(200)
	io.WriteString(w, `{"message": "success","data" :`)
	for _, val := range allusername {
		io.WriteString(w, "\""+val+"\",")
	}
	io.WriteString(w, "}")
}
