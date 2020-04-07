package MODELS

type SERVICES struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	IdUnit      int     `json:"idUnit"`
	Description string  `json:"description"`
}
