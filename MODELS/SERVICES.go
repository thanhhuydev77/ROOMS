package MODELS

type SERVICES struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	IdUnit      int     `json:"idUnit"`
	Description string  `json:"description"`
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

