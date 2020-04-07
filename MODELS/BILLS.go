package MODELS

import "time"

type BILLS struct {
	Id           int       `json:"id"`
	IdRoom       int       `json:"idRoom"`
	DateCheckOut time.Time `json:"dateCheckOut"`
	TotalPrice   float64   `json:"totalPrice"`
	IsCheckedOut int       `json:"isCheckedOut"`
}
