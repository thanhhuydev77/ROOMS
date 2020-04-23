package MODELS

import "ROOMS/COMMON"

type SERVICES struct {
	Id          int                 `json:"id"`
	NameService string              `json:"nameService"`
	Price       float64             `json:"price"`
	IdUnit      int                 `json:"idUnit"`
	Description COMMON.MyNullString `json:"description"`
	IdBlock     int                 `json:"idBlock"`
}

type GET_SERVICES_REQUEST struct {
	Id          int                 `json:"id"`
	NameService string              `json:"nameService"`
	Price       float64             `json:"price"`
	IdUnit      int                 `json:"idUnit"`
	Description COMMON.MyNullString `json:"description"`
	IdBlock     int                 `json:"idBlock"`
	UnitName    COMMON.MyNullString `json:"nameUnit"`
}

type SERVICE_INPUT struct {
	NameService string              `json:"nameService"`
	Price       float64             `json:"price"`
	IdUnit      int                 `json:"idUnit"`
	Description string `json:"description"`
	IdBlock     int                 `json:"idBlock"`
}

type SERVICES_INPUT struct {
	Services []SERVICE_INPUT		`json:"services"`
}


