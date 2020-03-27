package MODELS

import (
	"time"
)

type USERS struct {
	Id            int
	UserName      string
	Pass          string
	FullName      string
	IdentifyFront string
	IdentifyBack  string
	DateBirth     time.Time
	Address       string
	Role          int
	Sex           string
	Job           string
	WorkPlace     string
	TempReg       int
	Province      string
	Email         string
}
