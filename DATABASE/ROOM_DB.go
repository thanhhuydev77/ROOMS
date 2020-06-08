package DATABASE

import (
	"ROOMS/MODELS"
	"database/sql"
	"log"
	"strings"
)

func GetRoom(idBlock int) ([]MODELS.ROOMS, bool, error) {

	db, err := connectdatabase()

	if err != nil {
		log.Fatal("Can't connet to database")
		return nil, false, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM ROOMS WHERE idBlock = ?`, idBlock)
	if err != nil {
		log.Fatal(err)
		return nil, false, err
	}

	var rooms []MODELS.ROOMS

	for rows.Next() {
		var room MODELS.ROOMS
		err := rows.Scan(&room.Id, &room.Name, &room.Floor, &room.Square, &room.Price, &room.Description, &room.IdBlock, &room.MaxPeople, &room.Status, &room.CodeRoom)

		if err != nil {
			log.Fatal(err)
		}
		rooms = append(rooms, room)
	}
	return rooms, true, nil
}

func CreateRoom(room MODELS.ROOMS) (bool, error) {

	db, err := connectdatabase()

	if err != nil {
		log.Fatal("Cannot connect to database")
		return false, err
	}
	defer db.Close()

	row, err := db.Query(`INSERT INTO ROOMS(nameRoom, maxPeople,floor, square, price, description, idBlock, status) 
													VALUES(?,?,?,?,?,?,?,?)`,
		room.Name, room.MaxPeople, room.Floor, room.Square, room.Price,
		room.Description, room.IdBlock, room.Status)

	if err != nil {
		return false, err
	}
	defer row.Close()

	return true, nil
}

func DeleteRoom(id int) (bool, error) {

	db, err := connectdatabase()

	if err != nil {
		log.Fatal("Can't connect to database")
	}
	defer db.Close()

	row, err := db.Query(`DELETE FROM ROOMS WHERE id = ?`, id)

	if err != nil {
		return false, err
	}
	defer row.Close()

	return true, nil
}

func DeleteRooms(roomsId []int) (bool, error) {

	db, err := connectdatabase()

	if err != nil {
		return false, err
	}
	defer db.Close()

	args := make([]interface{}, len(roomsId))
	for i, id := range roomsId {
		args[i] = id
	}

	stmt := `DELETE FROM ROOMS WHERE id IN (?` + strings.Repeat(",?", len(args)-1) + `)`
	rows, err := db.Exec(stmt, args...)

	num, err := rows.RowsAffected()
	n := int64(num)

	if n == 0 {
		return false, err
	}

	return true, nil
}

func UpdateRoom(id int, room MODELS.ROOMS) (bool, error) {

	db, err := connectdatabase()

	if err != nil {
		return false, err
	}
	defer db.Close()

	row, err := db.Query(`UPDATE ROOMS 
										SET nameRoom = ? , maxPeople = ?, floor = ?, square = ? , price = ?, description = ?
										WHERE id = ?`,
		room.Name, room.MaxPeople, room.Floor, room.Square, room.Price, room.Description, id)

	if err != nil {
		return false, err
	}
	defer row.Close()

	return true, nil
}

func UpdateGetRoom(idBlock int) ([]MODELS.ROOMS, bool, error) {
	db, err := connectdatabase()

	if err != nil {
		log.Fatal("Can't connect to database")
		return nil, false, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT R.id, R.nameRoom, R.floor, R.square, R.price, R.description, R.idBlock, R.maxPeople	, R.status, R.codeRoom
FROM ROOMS R LEFT JOIN USER_ROOM UR ON R.id = UR.idRoom WHERE UR.idRoom IS NULL AND idBlock = ?`, idBlock)
	if err != nil {
		log.Fatal(err)
		return nil, false, err
	}

	var rooms []MODELS.ROOMS

	for rows.Next() {
		var room MODELS.ROOMS
		err := rows.Scan(&room.Id, &room.Name, &room.Floor, &room.Square, &room.Price, &room.Description, &room.IdBlock, &room.MaxPeople, &room.Status, &room.CodeRoom)

		if err != nil {
			log.Fatal(err)
		}
		rooms = append(rooms, room)
	}
	return rooms, true, nil
}

func GetRoomDB(idBlock int, status int, userid int) ([]MODELS.GET_ROOMDB_REQUEST, error) {
	db, err := connectdatabase()

	if err != nil {
		log.Fatal("Can't connect to database")
		return nil, err
	}
	defer db.Close()
	var rows *sql.Rows
	var err1 error
	if idBlock == -1 && status == -1 {
		rows, err1 = db.Query(`SELECT R.*, B.nameBlock, C.startDate FROM ROOMS R INNER JOIN BLOCKS B ON R.idBlock = B.id LEFT JOIN CONTRACTS C ON R.id = C.idRoom WHERE B.idOwner = ?`, userid)
	} else if idBlock == -1 && status != -1 {
		rows, err1 = db.Query(`SELECT R.*, B.nameBlock, C.startDate FROM ROOMS R INNER JOIN BLOCKS B ON R.idBlock = B.id LEFT JOIN CONTRACTS C ON R.id = C.idRoom WHERE B.idOwner = ? AND R.status = ?`, userid, status)
	} else {
		if status == -1 {
			rows, err1 = db.Query(`SELECT R.*, B.nameBlock, C.startDate FROM ROOMS R INNER JOIN BLOCKS B ON R.idBlock = B.id LEFT JOIN CONTRACTS C ON R.id = C.idRoom WHERE R.idBlock = ?`, idBlock)
		} else {
			rows, err1 = db.Query(`SELECT R.*, B.nameBlock, C.startDate FROM ROOMS R INNER JOIN BLOCKS B ON R.idBlock = B.id LEFT JOIN CONTRACTS C ON R.id = C.idRoom WHERE status = ? AND R.idBlock = ?`, status, idBlock)
		}
	}

	if err1 != nil {
		log.Fatal(err)
		return nil, err
	}

	var rooms []MODELS.GET_ROOMDB_REQUEST

	for rows.Next() {
		var room MODELS.GET_ROOMDB_REQUEST
		err := rows.Scan(&room.Id, &room.Name, &room.Floor, &room.Square, &room.Price, &room.Description, &room.IdBlock, &room.MaxPeople, &room.Status, &room.CodeRoom, &room.NameBlock, &room.StartDate)

		if err != nil {
			log.Fatal(err)
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func GetRoomImage(codeRoom string) ([]MODELS.ROOM_IMAGE, bool, error) {

	db, err := connectdatabase()

	if err != nil {
		log.Fatal("Can't connet to database")
		return nil, false, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM ROOM_IMAGES WHERE codeRoom = ?`, codeRoom)
	if err != nil {
		log.Fatal(err)
		return nil, false, err
	}

	var rooms []MODELS.ROOM_IMAGE

	for rows.Next() {
		var room MODELS.ROOM_IMAGE
		err := rows.Scan(&room.Id, &room.Name, &room.Status, &room.Url, &room.CodeRoom)

		if err != nil {
			log.Fatal(err)
		}
		rooms = append(rooms, room)
	}
	return rooms, true, nil
}

func GetUserRenting(codeRoom int) ([]MODELS.ROOM_USER_RENTING_NAME, bool, error) {

	db, err := connectdatabase()

	if err != nil {
		log.Fatal("Can't connet to database")
		return nil, false, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT C.fullName FROM USER_ROOM UR INNER JOIN CUSTOMERS C ON UR.idUser = C.id  WHERE UR.idRoom = ?`, codeRoom)
	if err != nil {
		log.Fatal(err)
		return nil, false, err
	}

	var rooms []MODELS.ROOM_USER_RENTING_NAME

	for rows.Next() {
		var room MODELS.ROOM_USER_RENTING_NAME
		err := rows.Scan(&room.Name)

		if err != nil {
			log.Fatal(err)
		}
		rooms = append(rooms, room)
	}
	return rooms, true, nil
}

func GetRoomById(id int) (*MODELS.ROOMS, error) {
	db, err := connectdatabase()

	if err != nil {
		log.Fatal("Can't connect to database")
		return nil, err
	}
	defer db.Close()
	var rows *sql.Rows
	var err1 error

	rows, err1 = db.Query(`SELECT * FROM ROOMS WHERE id = ?`, id)

	if err1 != nil {
		return nil, err
	}
	for rows.Next() {
		var room MODELS.ROOMS
		err := rows.Scan(&room.Id, &room.Name, &room.Floor, &room.Square, &room.Price, &room.Description, &room.IdBlock, &room.MaxPeople, &room.Status, &room.CodeRoom)

		if err != nil {
			return nil, err
		}
		return &room, nil
	}
	return nil, nil
}
