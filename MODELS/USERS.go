package MODELS

import (
	"database/sql"
	"fmt"
)

type USERS struct {
	Id            int
	UserName      string
	Pass          string
	FullName      sql.NullString
	IdentifyFront sql.NullString
	IdentifyBack  sql.NullString
	DateBirth     sql.NullTime
	Address       sql.NullString
	Role          sql.NullInt32
	Sex           sql.NullString
	Job           sql.NullString
	WorkPlace     sql.NullString
	TempReg       sql.NullInt32
	Province      sql.NullString
	Email         sql.NullString
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
		user.Role.Int32, user.Sex.String, user.Job.String, user.WorkPlace.String, user.TempReg.Int32, user.Province.String, user.Email.String)
}
