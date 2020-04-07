package MODELS

import (
	. "ROOMS/COMMON"
	"fmt"
)

//this type use to save data requested from database
type USERS struct {
	Id            int          `json:"id"`
	UserName      string       `json:"userName"`
	Pass          string       `json:"pass"`
	FullName      MyNullString `json:"fullName"`
	IdentifyFront MyNullString `json:"identifyFront"`
	IdentifyBack  MyNullString `json:"identifyBack"`
	DateBirth     MyNullTime   `json:"dateBirth"`
	Address       MyNullString `json:"address"`
	Role          MyNullInt    `json:"role"`
	Sex           MyNullString `json:"sex"`
	Job           MyNullString `json:"job"`
	WorkPlace     MyNullString `json:"workPlace"`
	TempReg       MyNullInt    `json:"tempReg"`
	Province      MyNullString `json:"province"`
	Email         MyNullString `json:"email"`
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
