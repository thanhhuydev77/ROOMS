package MODELS

type EQUIPMENTS struct {
	Id          int     `json:"id"`
	IdRoom      int     `json:"idRoom"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	ComPrice    float64 `json:"comPrice"`
	IdUnit      int     `json:"idUnit"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
}
