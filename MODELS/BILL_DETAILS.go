package MODELS

type BILL_DETAILS struct {
	Id         int     `json:"id"`
	IdBill     int     `json:"idBill"`
	IdService  int     `json:"idService"`
	Amount     int     `json:"amount"`
	TotalPrice float64 `json:"totalPrice"`
}
