package DATABASE

import (
	"ROOMS/MODELS"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func GetCustomers(db *sql.DB, userId int) ([]MODELS.CUSTOMER, bool, error) {
	var listCustomers []MODELS.CUSTOMER
	//db, err := connectdatabase()
	//
	//if err != nil {
	//	return nil, false, err
	//}
	//defer db.Close()
	if db == nil {
		log.Print("can not connect to database!")
		return nil, false, fmt.Errorf("can not connect db")
	}

	rows, err := db.Query("SELECT CU.*, R.nameRoom FROM CUSTOMERS CU LEFT JOIN USER_ROOM UR ON CU.id = UR.idUser "+
		"LEFT JOIN ROOMS R ON UR.idRoom = R.id  WHERE idOwner = ?", userId)
	if err != nil {
		log.Fatal(err)
		return nil, false, err
	}

	for rows.Next() {
		var c MODELS.CUSTOMER

		err := rows.Scan(&c.Id, &c.CodeUser, &c.UserName, &c.Pass, &c.FullName,
			&c.IdentifyFront, &c.IdentifyBack, &c.DateBirth, &c.Address,
			&c.Role, &c.Sex, &c.Job, &c.WorkPlace, &c.TempReg, &c.Province,
			&c.Email, &c.Avatar, &c.PhoneNumber, &c.IdOwner, &c.Note, &c.NameRoom)

		if err != nil {
			log.Fatal(err)
			return nil, false, err
		}
		listCustomers = append(listCustomers, c)
	}
	defer rows.Close()
	return listCustomers, true, nil
}

func GetCustomersByUserId(db *sql.DB, userId int) ([]MODELS.CUSTOMER_GET, bool, error) {
	var listCustomers []MODELS.CUSTOMER_GET
	//db, err := connectdatabase()
	//
	//if err != nil {
	//	return nil, false, err
	//}
	//defer db.Close
	if db == nil {
		log.Print("can not connect to database!")
		return nil, false, fmt.Errorf("can not connect db")
	}

	rows, err := db.Query("SELECT * FROM CUSTOMERS WHERE idOwner = ?", userId)
	if err != nil {
		log.Fatal(err)
		return nil, false, err
	}

	for rows.Next() {
		var c MODELS.CUSTOMER_GET

		err := rows.Scan(&c.Id, &c.CodeUser, &c.UserName, &c.Pass, &c.FullName,
			&c.IdentifyFront, &c.IdentifyBack, &c.DateBirth, &c.Address,
			&c.Role, &c.Sex, &c.Job, &c.WorkPlace, &c.TempReg, &c.Province,
			&c.Email, &c.Avatar, &c.PhoneNumber, &c.IdOwner, &c.Note)

		var r, _, _ = SelectNameRoom(db, c.Id)
		for i := range r {
			c.Rooms = append(c.Rooms, MODELS.CUSTOMER_ROOMS{RoomName: r[i]})
		}

		if err != nil {
			log.Fatal(err)
			return nil, false, err
		}
		listCustomers = append(listCustomers, c)
	}
	defer rows.Close()
	return listCustomers, true, nil
}

func CreateCustomer(db *sql.DB, c MODELS.CUSTOMER_INPUT) (bool, error) {
	//db, err := connectdatabase()
	//if err != nil {
	//	log.Fatalln(err)
	//	return false, err
	//}
	//defer db.Close()
	if db == nil {
		log.Print("can not connect to database!")
		return false, fmt.Errorf("can not connect db")
	}
	rs, err := db.Query(`INSERT INTO CUSTOMERS (codeUser, fullName, identifyFront, identifyBack, dateBirth, sex, 
		job, workPlace, tempReg, email, avatar, phoneNumber, idOwner, note) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		c.CodeUser, c.FullName, c.IdentifyFront, c.IdentifyBack, c.DateBirth, c.Sex, c.Job,
		c.WorkPlace, c.TempReg, c.Email, c.Avatar, c.PhoneNumber, c.IdOwner, c.Note)

	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	defer rs.Close()

	return true, nil
}

func DeleteCustomer(db *sql.DB, idCustomer int) (bool, error) {
	//db, err := connectdatabase()
	//if err != nil {
	//	log.Fatalln(err)
	//	return false, err
	//}
	//defer db.Close()
	if db == nil {
		log.Print("can not connect to database!")
		return false, fmt.Errorf("can not connect db")
	}
	rs, err := db.Exec(`DELETE FROM CUSTOMERS WHERE id = ?`, idCustomer)

	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	num, err := rs.RowsAffected()
	if num == 0 {
		return false, err
	}
	return true, nil
}

func DeleteManyCustomers(db *sql.DB, ids []int) (bool, error) {
	//db, err := connectdatabase()
	//if err != nil {
	//	log.Fatalln(err)
	//	return false, err
	//}
	//defer db.Close()
	if db == nil {
		log.Print("can not connect to database!")
		return false, fmt.Errorf("can not connect db")
	}
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	stmt := `DELETE FROM CUSTOMERS WHERE id IN(?` + strings.Repeat(",?", len(args)-1) + `)`
	rows, err := db.Exec(stmt, args...)

	num, err := rows.RowsAffected()
	m := int64(num)
	if m == 0 {
		return false, err
	}
	return true, nil
}

func UpdateCustomer(db *sql.DB, c MODELS.CUSTOMER_UPDATE) (bool, error) {
	//db, err := connectdatabase()
	//
	//if err != nil {
	//
	//	log.Print("can not connect to database!")
	//	return false, err
	//}
	//defer db.Close()
	if db == nil {
		log.Print("can not connect to database!")
		return false, fmt.Errorf("can not connect db")
	}

	rows, err := db.Exec(`UPDATE CUSTOMERS SET fullName = ?, identifyFront = ?, identifyBack = ?, dateBirth = ?, 
address = ?, sex = ?, job = ?, workPlace = ?, tempReg = ?, email = ?, avatar = ?, phoneNumber = ?, note = ? WHERE id = ?`,
		c.FullName, c.IdentifyFront, c.IdentifyBack, c.DateBirth, c.Address, c.Sex, c.Job,
		c.WorkPlace, c.TempReg, c.Email, c.Avatar, c.PhoneNumber, c.Note, c.Id)

	num, err := rows.RowsAffected()
	m := int64(num)
	if m == 0 {
		return false, err
	}
	return true, nil
}

func SelectNameRoom(db *sql.DB, idUser int) ([]string, bool, error) {
	//db, err := connectdatabase()
	//
	//if err != nil {
	//	log.Fatal("Can't connet to database")
	//	return nil, false, err
	//}
	//defer db.Close()
	if db == nil {
		log.Print("can not connect to database!")
		return nil, false, fmt.Errorf("can not connect db")
	}

	rows, err := db.Query(`SELECT R.nameRoom FROM USER_ROOM UR INNER JOIN ROOMS R ON UR.idRoom = R.id WHERE idUser = ?`, idUser)
	if err != nil {
		log.Fatal(err)
		return nil, false, err
	}

	var rooms []string

	for rows.Next() {
		var room string
		err := rows.Scan(&room)

		if err != nil {
			log.Fatal(err)
		}
		rooms = append(rooms, room)
	}
	return rooms, true, nil
}
