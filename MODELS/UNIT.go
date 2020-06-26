package MODELS

import (
	"ROOMS/COMMON"
)

type UNIT struct {
	Id          int                 `json:"id"`
	Name        string              `json:"name"`
	Description COMMON.MyNullString `json:"description"`
}
