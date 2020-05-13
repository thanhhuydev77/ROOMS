package MODELS

import (
	"ROOMS/COMMON"
	"time"
)

type CUSTOMER struct {
	Id	int       	`json:"id"`
	CodeUser string `json:"codeUser"`
	UserName COMMON.MyNullString `json:"userName"`
	Pass COMMON.MyNullString `json:"pass"`
	FullName string `json:"fullName"`
	IdentifyFront string `json:"identifyFront"`
	IdentifyBack string `json:"identifyBack"`
	DateBirth time.Time `json:"dateBirth"`
	Address COMMON.MyNullString `json:"address"`
	Role COMMON.MyNullInt `json:"role"`
	Sex string `json:"sex"`
	Job string `json:"job"`
	WorkPlace string `json:"workPlace"`
	TempReg int `json:"tempReg"`
	Province COMMON.MyNullString `json:"province"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	PhoneNumber string `json:"phoneNumber"`
	IdOwner int `json:"idOwner"`
	Note string `json:"note"`
	Status int `json:"status"`
	NameRoom string `json:"nameRoom"`
}

type CUSTOMER_INPUT struct {
	FullName string 	`json:"fullName"`
	PhoneNumber string 	`json:"phoneNumber"`
	DateBirth string 	`json:"dateBirth"`
	Email string 		`json:"email"`
	Job string 			`json:"job"`
	WorkPlace string	`json:"workPlace"`
	Sex string 			`json:"sex"`
	TempReg int 		`json:"tempReg"`
	Note string 		`json:"note"`
	Avatar string 		`json:"avatar"`
	IdentifyBack string `json:"identifyBack"`
	IdentifyFront string `json:"identifyFront"`
	CodeUser string		`json:"codeUser"`
	IdOwner int 		`json:"idOwner"`
}