package MODELS

import "ROOMS/COMMON"

type ROOMS struct {
	Id          int                 `json:"id"`
	Name        string              `json:"nameRoom"`
	Floor       int                 `json:"floor"`
	Square      int                 `json:"square"`
	Price       float64             `json:"price"`
	Description string              `json:"description"`
	IdBlock     int                 `json:"idBlock"`
	MaxPeople   int                 `json:"maxPeople"`
	Status      int                 `json:"status"`
	CodeRoom    COMMON.MyNullString `json:"codeRoom"`
}

type ROOMIDS struct {
	RoomsId []int
}

type GET_ROOMDB_REQUEST struct {
	Id          int                 `json:"id"`
	Name        string              `json:"nameRoom"`
	Floor       int                 `json:"floor"`
	Square      int                 `json:"square"`
	Price       float64             `json:"price"`
	Description string              `json:"description"`
	IdBlock     int                 `json:"idBlock"`
	MaxPeople   int                 `json:"maxPeople"`
	Status      int                 `json:"status"`
	CodeRoom    COMMON.MyNullString `json:"codeRoom"`
	NameBlock   string              `json:"nameBlock"`
	StartDate   COMMON.MyNullString `json:"startDate"`
}

type ROOM_IMAGE struct {
	Id       int                 `json:"id"`
	Name     COMMON.MyNullString `json:"name"`
	Status   COMMON.MyNullString `json:"status"`
	Url      COMMON.MyNullString `json:"url"`
	CodeRoom COMMON.MyNullString `json:"codeRoom"`
}
