package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
)

//get all room in block with idblock
func GetRoom(db *sql.DB, idBlock int) ([]MODELS.ROOMS, bool, error) {
	return DATABASE.GetRoom(db, idBlock)
}

//create a new room
func CreateRoom(db *sql.DB, room MODELS.ROOMS) (bool, error) {
	return DATABASE.CreateRoom(db, room)
}

//delte a room with its id
func DeleteRoom(db *sql.DB, id int) (bool, error) {
	return DATABASE.DeleteRoom(db, id)
}

//delete many room
func DeleteRooms(db *sql.DB, roomsId []int) (bool, error) {
	return DATABASE.DeleteRooms(db, roomsId)
}

//update a room
func UpdateRoom(db *sql.DB, id int, room MODELS.ROOMS) (bool, error) {
	return DATABASE.UpdateRoom(db, id, room)
}

//get all room in block with idblock
func UpdateGetRoom(db *sql.DB, idBlock int) ([]MODELS.ROOMS, bool, error) {
	return DATABASE.UpdateGetRoom(db, idBlock)
}

//get room dashboard
func GetRoomDB(db *sql.DB, idBlock int, status int, userid int) ([]MODELS.GET_ROOMDB_REQUEST, error) {
	return DATABASE.GetRoomDB(db, idBlock, status, userid)
}

//get room image
func GetRoomImage(db *sql.DB, codeRoom string) ([]MODELS.ROOM_IMAGE, bool, error) {
	return DATABASE.GetRoomImage(db, codeRoom)
}

//get user renting
func GetUserRenting(db *sql.DB, codeRoom int) ([]MODELS.ROOM_USER_RENTING_NAME, bool, error) {
	return DATABASE.GetUserRenting(db, codeRoom)
}

//get room by id
func GetRoomById(db *sql.DB, id int) (*MODELS.ROOMS, error) {
	return DATABASE.GetRoomById(db, id)
}
