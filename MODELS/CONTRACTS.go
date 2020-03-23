package MODELS

import "time"

type CONTRACTS struct {
	id        int
	idRoom    int
	idOwner   int
	idSlave   int
	startDate time.Time
	endDate   time.Time
	circlePay int
}
