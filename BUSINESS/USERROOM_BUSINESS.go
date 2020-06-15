package BUSINESS

import "ROOMS/DATABASE"

func DeleteManyUserRoom(ids []int) (bool, error)  {
	return DATABASE.DeleteManyUserRoom(ids)
}