package MODELS

import (
	"ROOMS/COMMON"
)

type BILLS struct {
	Id           int               `json:"id"`
	IdRoom       int               `json:"idRoom"`
	DateCheckOut COMMON.MyNullTime `json:"dateCheckOut"`
	TotalPrice   float64           `json:"totalPrice"`
	IsCheckedOut int               `json:"isCheckedOut"`
}

type CREATE_UPDATE_BILL_REQUEST struct {
	Id           int            `json:"id"`
	IdRoom       int            `json:"idRoom"`
	DateCheckOut string         `json:"dateCheckOut"`
	TotalPrice   float64        `json:"totalPrice"`
	IsCheckedOut int            `json:"isCheckedOut"`
	BillDetail   []BILL_DETAILS `json:"billDetail"`
}
