package BUSINESS

import (
	"ROOMS/MODELS"
	"ROOMS/STATICS"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

func Login(username string, pass string) (bool, bool, MODELS.USERS) {
	exsist := false
	passOK := false
	a := MODELS.USERS{}
	//getuset from datebase
	db, err := STATICS.Connectdatabase()
	// Query all users
	if db == nil {

		log.Print("can not connect to database!")
		return exsist, false, a
	}
	defer db.Close()

	rows, err := db.Query("select id,username,pass,fullname,role from USERS where username = ?", username)
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
	//hashpass,err := HashPassword(pass)
	if CheckPasswordHash(pass, a.Pass) {
		passOK = true
	}
	return exsist, passOK, a
}

func Register(user MODELS.RequestRegister) (bool, error) {

	//getuset from datebase
	db, err := STATICS.Connectdatabase()
	// Query all users
	if err != nil {

		log.Print("can not connect to database!")
		return false, err
	}
	defer db.Close()
	passhash, _ := HashPassword(user.Pass)
	rows, err := db.Query(`insert into USERS(userName,Pass,FullName,Address,Role,Sex,Province,Email)
							  values(?,?,?,?,?,?,?,?)`, user.UserName, passhash, user.FullName, user.Address, user.Role, user.Sex, user.Province, user.Email)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return true, nil
}

func GetAllUserName() []string {
	var Allusername []string
	db, err := STATICS.Connectdatabase()
	// Query all users
	if db == nil {

		log.Print("can not connect to database!")
		return nil
	}
	defer db.Close()

	rows, err := db.Query("select username from USERS")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
		Allusername = append(Allusername, username)
	}
	defer rows.Close()
	return Allusername
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUsers(Id int) []MODELS.USERS {

	db, err := STATICS.Connectdatabase()
	// Query all users
	if db == nil {

		log.Print("can not connect to database!")
		return nil
	}
	defer db.Close()
	query := ""
	list := []MODELS.USERS{}

	if Id == -1 {
		query = "select * from USERS"
	} else {
		query = "select * from USERS where Id = " + strconv.Itoa(Id)
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user MODELS.USERS
		err := rows.Scan(&user.Id, &user.UserName, &user.Pass, &user.FullName, &user.IdentifyFront, &user.IdentifyBack, &user.DateBirth, &user.Address,
			&user.Role, &user.Sex, &user.Job, &user.WorkPlace, &user.TempReg, &user.Province, &user.Email, &user.Avatar, &user.PhoneNumber)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, user)
	}
	return list
}
