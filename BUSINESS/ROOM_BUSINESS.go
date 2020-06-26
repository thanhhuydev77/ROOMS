package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
)

func GetRoom(db *sql.DB, idBlock int) ([]MODELS.ROOMS, bool, error) {
	return DATABASE.GetRoom(db, idBlock)
}

func CreateRoom(db *sql.DB, room MODELS.ROOMS) (bool, error) {
	return DATABASE.CreateRoom(db, room)
}

func DeleteRoom(db *sql.DB, id int) (bool, error) {
	return DATABASE.DeleteRoom(db, id)
}

func DeleteRooms(db *sql.DB, roomsId []int) (bool, error) {
	return DATABASE.DeleteRooms(db, roomsId)
}

func UpdateRoom(db *sql.DB, id int, room MODELS.ROOMS) (bool, error) {
	return DATABASE.UpdateRoom(db, id, room)
}

func UpdateGetRoom(db *sql.DB, idBlock int) ([]MODELS.ROOMS, bool, error) {
	return DATABASE.UpdateGetRoom(db, idBlock)
}

func GetRoomDB(db *sql.DB, idBlock int, status int, userid int) ([]MODELS.GET_ROOMDB_REQUEST, error) {
	return DATABASE.GetRoomDB(db, idBlock, status, userid)
}

func GetRoomImage(db *sql.DB, codeRoom string) ([]MODELS.ROOM_IMAGE, bool, error) {
	return DATABASE.GetRoomImage(db, codeRoom)
}

func GetUserRenting(db *sql.DB, codeRoom int) ([]MODELS.ROOM_USER_RENTING_NAME, bool, error) {
	return DATABASE.GetUserRenting(db, codeRoom)
}

func GetRoomById(db *sql.DB, id int) (*MODELS.ROOMS, error) {
	return DATABASE.GetRoomById(db, id)
}
