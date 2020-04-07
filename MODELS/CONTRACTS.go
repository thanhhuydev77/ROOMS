package MODELS

import "time"

type CONTRACTS struct {
	Id        int       `json:"id"`
	IdRoom    int       `json:"idRoom"`
	IdOwner   int       `json:"idOwner"`
	IdSlave   int       `json:"idSlave"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	CirclePay int       `json:"circlePay"`
}
