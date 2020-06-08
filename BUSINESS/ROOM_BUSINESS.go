package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
)

func GetRoom(idBlock int) ([]MODELS.ROOMS, bool, error) {
	return DATABASE.GetRoom(idBlock)
}

func CreateRoom(room MODELS.ROOMS) (bool, error) {
	return DATABASE.CreateRoom(room)
}

func DeleteRoom(id int) (bool, error) {
	return DATABASE.DeleteRoom(id)
}

func DeleteRooms(roomsId []int) (bool, error) {
	return DATABASE.DeleteRooms(roomsId)
}

func UpdateRoom(id int, room MODELS.ROOMS) (bool, error) {
	return DATABASE.UpdateRoom(id, room)
}

func UpdateGetRoom(idBlock int) ([]MODELS.ROOMS, bool, error) {
	return DATABASE.UpdateGetRoom(idBlock)
}

func GetRoomDB(idBlock int, status int, userid int) ([]MODELS.GET_ROOMDB_REQUEST, error) {
	return DATABASE.GetRoomDB(idBlock, status, userid)
}

func GetRoomImage(codeRoom string) ([]MODELS.ROOM_IMAGE, bool, error) {
	return DATABASE.GetRoomImage(codeRoom)
}

func GetUserRenting(codeRoom int) ([]MODELS.ROOM_USER_RENTING_NAME, bool, error) {
	return DATABASE.GetUserRenting(codeRoom)
}

func GetRoomById(id int) (*MODELS.ROOMS, error) {
	return DATABASE.GetRoomById(id)
}
