package MODELS

import "time"

type BILLS struct {
	Id           int
	IdRoom       int
	DateCheckOut time.Time
	TotalPrice   float64
	IsCheckedOut int
}
