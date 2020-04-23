package MODELS

import "ROOMS/COMMON"

type DEFAULT_SERVICES struct {
	Id          int     `json:"id"`
	NameService string  `json:"nameService"`
	Price       float64 `json:"price"`
	IdUnit      int     `json:"idUnit"`
	Description string  `json:"description"`
}
type GET_DEFAULT_SERVICES_REQUEST struct {
	Id          int                 `json:"id"`
	NameService string              `json:"nameService"`
	Price       float64             `json:"price"`
	Description COMMON.MyNullString `json:"description"`
	IdUnit      int                 `json:"idUnit"`
	UnitName    COMMON.MyNullString `json:"name"`
}
