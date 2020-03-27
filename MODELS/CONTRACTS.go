package MODELS

import "time"

type CONTRACTS struct {
	Id        int
	IdRoom    int
	IdOwner   int
	IdSlave   int
	StartDate time.Time
	EndDate   time.Time
	CirclePay int
}
