package BUSINESS

import (
	"ROOMS/MODELS"
	"ROOMS/STATICS"
	"log"
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
	if pass == a.Pass {
		passOK = true
	}
	return exsist, passOK, a
}

func Register(user MODELS.USERS) bool {

	//getuset from datebase
	db, err := STATICS.Connectdatabase()
	// Query all users
	if db == nil {

		log.Print("can not connect to database!")
		return false
	}
	defer db.Close()

	rows, err := db.Query(`insert into USERS(userName,Pass,FullName,IdentifyFront,IdentifyBack
							  ,DateBirth,Address,Role,Sex,Job,WorkPlace,TempReg,Province,Email)
							  values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, user.UserName, user.Pass, user.FullName, user.IdentifyFront, user.IdentifyBack,
		user.DateBirth, user.Address, user.Role, user.Sex, user.Job,
		user.WorkPlace, user.TempReg, user.Province, user.Email)
	if err != nil {
		return false
	}

	defer rows.Close()

	return true
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
