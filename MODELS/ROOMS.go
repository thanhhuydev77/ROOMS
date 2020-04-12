package MODELS

type ROOMS struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Floor       int     `json:"floor"`
	Square      int     `json:"square"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	IdBlock     int     `json:"idBlock"`
	MaxPeople   int     `json:"maxPeople"`
	Status 		int		`json:"status"`
}
