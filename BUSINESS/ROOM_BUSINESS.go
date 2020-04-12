package BUSINESS

import (
	"ROOMS/MODELS"
	"ROOMS/STATICS"
	"log"
	"strings"
)

func CreateRoom(room MODELS.ROOMS) (bool, error)  {

	db, err := STATICS.Connectdatabase()

	if err != nil {
		log.Fatal("Cannot connect to database")
		return false, err
	}
	defer db.Close()

	row, err := db.Query(`INSERT INTO ROOMS(nameRoom, maxPeople,floor, square, price, description, idBlock, status) 
													VALUES(?,?,?,?,?,?,?,?)`,
													room.Name, room.MaxPeople, room.Floor, room.Square, room.Price,
													room.Description, room.IdBlock, room.Status)

	if err != nil{
		return false, err
	}
	defer row.Close()

	return true, nil
}

func DeleteRoom(id int) (bool , error) {

	db, err := STATICS.Connectdatabase()

	if err != nil{
		log.Fatal("Can't connect to database")
	}
	defer db.Close()

	row, err := db.Query(`DELETE FROM ROOMS WHERE id = ?`, id)

	if err != nil{
		return false, err
	}
	defer row.Close()

	return true, nil
}

func DeleteRooms(roomsId []int)(bool, error)  {

	db, err := STATICS.Connectdatabase()

	if err != nil{
		return false, err
	}
	defer db.Close()

	args := make([]interface{}, len(roomsId))
	for i, id := range roomsId {
		args[i] = id
	}

	stmt := `DELETE FROM ROOMS WHERE id IN (?` + strings.Repeat(",?", len(args)-1) + `)`
	rows, err := db.Exec(stmt, args...)

	num , err := rows.RowsAffected()
	n := int64(num)

	if n == 0{
		return false, err
	}

	return true, nil
}

func UpdateRoom(id int, room MODELS.ROOMS) (bool, error)  {

	db, err := STATICS.Connectdatabase()

	if err != nil{
		return false, err
	}
	defer db.Close()

	row, err := db.Query(`UPDATE ROOMS 
										SET nameRoom = ? , maxPeople = ?, floor = ?, square = ? , price = ?, description = ?
										WHERE id = ?`,
										room.Name, room.MaxPeople, room.Floor, room.Square, room.Price, room.Description, id)

	if err != nil{
		return false, err
	}
	defer row.Close()

	return true, nil
}