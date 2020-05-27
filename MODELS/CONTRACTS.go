package MODELS

import (
	"ROOMS/COMMON"
	"time"
)

type CONTRACTS struct {
	Id        int                 `json:"id"`
	IdRoom    int                 `json:"idRoom"`
	IdOwner   int                 `json:"idOwner"`
	IdSlave   int                 `json:"idSlave"`
	StartDate time.Time           `json:"startDate"`
	EndDate   time.Time           `json:"endDate"`
	CirclePay int                 `json:"circlePay"`
	Deposit   float64             `json :"deposit"`
	DayPay    int                 `json:"dayPay"`
	Note      COMMON.MyNullString `json:"note"`
	IdBlock   int                 `json:"idBlock"`
}

type GET_CONTRACTS_REQUEST struct {
	Id        int                 `json:"id"`
	IdRoom    int                 `json:"idRoom"`
	IdOwner   int                 `json:"idOwner"`
	IdSlave   int                 `json:"idSlave"`
	StartDate COMMON.MyNullString `json:"startDate"`
	EndDate   COMMON.MyNullString `json:"endDate"`
	CirclePay int                 `json:"circlePay"`
	Deposit   float64             `json:"deposit"`
	DayPay    int                 `json:"dayPay"`
	Note      string              `json:"note"`
	IdBlock   int                 `json:"idBlock"`
	NameRoom  COMMON.MyNullString `json:"nameRoom"`
	FullName  COMMON.MyNullString `json:"fullName"`
	IdUsers   []int               `json:"idUsers"`
}

type CREATE_UPDATE_CONTRACT_REQUEST struct {
	Id        int         `json:"id"`
	IdRoom    int         `json:"idRoom"`
	IdOwner   int         `json:"idOwner"`
	IdSlave   int         `json:"idSlave"`
	StartDate string      `json:"startDate"`
	EndDate   string      `json:"endDate"`
	CirclePay int         `json:"circlePay"`
	Deposit   float64     `json :"deposit"`
	DayPay    int         `json:"dayPay"`
	Note      string      `json:"note"`
	IdBlock   int         `json:"idBlock"`
	UserRooms []USER_ROOM `json:"userRooms"`
}
