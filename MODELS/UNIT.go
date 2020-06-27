package MODELS

import "ROOMS/COMMON"

type UNIT struct {
	Id          int                 `json:"id"`
	Name        string              `json:"name"`
	Description COMMON.MyNullString `json:"description"`
}

type RespondOk struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type RespondFail struct {
	Message string `json:"message"`
}
