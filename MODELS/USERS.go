package MODELS

import (
	. "ROOMS/COMMON"
	"fmt"
)

//this type use to save data requested from database
type USERS struct {
	Id            int
	UserName      string
	Pass          string
	FullName      MyNullString
	IdentifyFront MyNullString
	IdentifyBack  MyNullString
	DateBirth     MyNullTime
	Address       MyNullString
	Role          MyNullInt
	Sex           MyNullString
	Job           MyNullString
	WorkPlace     MyNullString
	TempReg       MyNullInt
	Province      MyNullString
	Email         MyNullString
}

func (user USERS) String() string {
	return fmt.Sprintf(`{  
				"id": %v,
                "userName": "%v",
                "pass": "%v",
                "fullName": "%v",
                "identifyFront": "%v",
                "identifyBack": "%v",
                "dateBirth": "%v",
                "address": "%v",
                "role": %v,
                "sex": "%v",
                "job": "%v",
                "workPlace": "%v",
                "tempReg": %v,
                "province": %v,
                "email": %v
}`, user.Id, user.UserName, user.Pass, user.FullName.String, user.IdentifyFront.String, user.IdentifyBack.String, user.DateBirth.Time, user.Address.String,
		user.Role.Int64, user.Sex.String, user.Job.String, user.WorkPlace.String, user.TempReg.Int64, user.Province.String, user.Email.String)
}

//this type use to save request body of method register
type RequestRegister struct {
	UserName string
	Pass     string
	Confirm  string
	FullName string
	Sex      string
	Email    string
	Role     int
	Province string
	Address  string
}
