package CONTROLLERS

import (
<<<<<<< HEAD
	"ROOMS/MODELS"
	"ROOMS/STATICS"
	"github.com/dgrijalva/jwt-go"
=======
	"ROOMS/MIDDLEWARE"
	"ROOMS/MODELS"
	"ROOMS/STATICS"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
>>>>>>> origin/Huy
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

<<<<<<< HEAD
=======
func InitUserController(r *mux.Router) {
	r.HandleFunc("/User/login", TokenHandler).Methods("POST")
	r.Handle("/user/getall", MIDDLEWARE.AuthMiddleware(http.HandlerFunc(getalluser)))
}

>>>>>>> origin/Huy
func TokenHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	a := MODELS.USERS{}

	r.ParseForm()

	a.UserName = r.Form.Get("UserName")
	tempPass := r.Form.Get("Pass")
	//getuset from datebase
	db, err := STATICS.Connectdatabase()
	// Query all users
	if db == nil {
		log.Print("can not connect to database!")
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	defer db.Close()

	exsist := false
	rows, err := db.Query("select id,username,pass,fullname,role from USERS where username = ?", a.UserName)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&a.Id, &a.UserName, &a.Pass, &a.FullName, &a.Role)
		if err != nil {
			log.Fatal(err)
		}
		exsist = true
	}
	defer rows.Close()

	if !exsist {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"message": "Can't find user please sign in again!"}`)
		return
	}

	if exsist && tempPass != a.Pass {
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
	tokenString, err := token.SignedString([]byte(STATICS.APP_KEY))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"token_generation_failed"}`)
		return
	}
	a.Role = 1
	a.Id = 1
	//w.Write([]byte(`{"hello": "world"}`))
	stringresult := `{"message": "Login success","data":{"token":"` + tokenString + `","user":{ "id":` + strconv.Itoa(a.Id) + `,
						"username":"` + a.UserName + `","fullname":"` + a.FullName + `","role":` + strconv.Itoa(a.Role) + `}}}`
	io.WriteString(w, stringresult)
	return
}

func getalluser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"ok all user"}`)
}
