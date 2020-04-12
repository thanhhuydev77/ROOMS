package BUSINESS

import (
	"ROOMS/MODELS"
	"ROOMS/STATICS"
	"log"
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