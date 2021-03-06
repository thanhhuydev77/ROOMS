package DATABASE

import (
	"ROOMS/MODELS"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

//get contracts by blockid
func GetContractByBlockId(db *sql.DB, BlockId int) ([]MODELS.GET_CONTRACTS_REQUEST, bool, error) {
	var listContracts []MODELS.GET_CONTRACTS_REQUEST
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
	rows, err := db.Query("SELECT C.*, R.nameRoom, CU.fullName FROM CONTRACTS C INNER JOIN ROOMS R ON C.idRoom = R.id INNER JOIN CUSTOMERS CU ON C.idSlave = CU.id WHERE C.idBlock = ?", BlockId)
	if err != nil {
		log.Fatal(err)
		return nil, false, err
	}

	for rows.Next() {
		var c MODELS.GET_CONTRACTS_REQUEST

		err := rows.Scan(&c.Id, &c.IdRoom, &c.IdOwner, &c.IdSlave, &c.StartDate,
			&c.EndDate, &c.CirclePay, &c.Deposit, &c.DayPay,
			&c.Note, &c.IdBlock, &c.NameRoom, &c.FullName)
		row2, err := db.Query("SELECT idUser FROM USER_ROOM WHERE idRoom = ?", c.IdRoom)
		if err != nil {
			log.Fatal(err)
			return nil, false, err
		}
		for row2.Next() {
			var a int
			err := row2.Scan(&a)
			c.IdUsers = append(c.IdUsers, a)

			if err != nil {
				log.Fatal(err)
				return nil, false, err
			}
		}

		if err != nil {
			log.Fatal(err)
			return nil, false, err
		}
		listContracts = append(listContracts, c)
	}
	defer rows.Close()
	return listContracts, true, nil
}

//create a new contract
func CreateContract(db *sql.DB, CCR MODELS.CREATE_UPDATE_CONTRACT_REQUEST) bool {
	//db, err := connectdatabase()
	//
	//if err != nil {
	//	log.Fatal("Can't connet to database")
	//	return false
	//}
	//defer db.Close()
	if db == nil {
		log.Print("can not connect to database!")
		return false
	}
	// Query all users

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO CONTRACTS (idRoom, idOwner, idSlave, startDate, endDate, circlePay, deposit, dayPay, note, idBlock) VALUES (?,?,?,?,?,?,?,?,?,?)`, CCR.IdRoom, CCR.IdOwner, CCR.IdSlave, CCR.StartDate, CCR.EndDate, CCR.CirclePay, CCR.Deposit, CCR.DayPay, CCR.Note, CCR.IdBlock)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return false
	}

	// The next query is handled similarly
	for _, vals := range CCR.UserRooms {
		_, err = tx.ExecContext(ctx, `INSERT INTO USER_ROOM (idUser, idRoom) VALUES(?,?)`, vals.IdUser, vals.IdRoom)
		if err != nil {
			// Incase we find any error in the query execution, rollback the transaction
			tx.Rollback()
			return false
		}
	}

	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = nil
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	//defer rows.Close()
	return true
}

//delete a contract
func DeleteContract(db *sql.DB, idCustomer int) (bool, error) {
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
	rs, err := db.Exec(`DELETE FROM CONTRACTS WHERE id = ?`, idCustomer)

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

//delete many contract
func DeleteAllContract(db *sql.DB, idContract []int) (bool, error) {
	//db, err := connectdatabase()
	//
	//if err != nil {
	//	return false, err
	//}
	//defer db.Close()
	if db == nil {
		log.Print("can not connect to database!")
		return false, fmt.Errorf("can not connect db")
	}
	args := make([]interface{}, len(idContract))
	for i, id := range idContract {
		args[i] = id
	}

	stmt := `DELETE FROM CONTRACTS WHERE id IN (?` + strings.Repeat(",?", len(args)-1) + `)`
	rows, err := db.Exec(stmt, args...)

	num, err := rows.RowsAffected()
	n := int64(num)

	if n == 0 {
		return false, err
	}

	return true, nil
}

//update a contract
func UpdateContract(db *sql.DB, c MODELS.CREATE_UPDATE_CONTRACT_REQUEST) (bool, error) {
	//db, err := connectdatabase()
	//if err != nil {
	//	log.Print("can not connect to database!")
	//	return false, err
	//}
	//defer db.Close()
	if db == nil {
		log.Print("can not connect to database!")
		return false, fmt.Errorf("can not connect db")
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("UPDATE CONTRACTS SET idSlave = ?, startDate = ?, endDate = ?, circlePay = ?, deposit = ?, dayPay = ?, note = ? WHERE id = ?", c.IdSlave, c.StartDate, c.EndDate, c.CirclePay, c.Deposit, c.DayPay, c.Note, c.Id)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return false, err
	}

	_, err2 := db.Exec("DELETE FROM USER_ROOM WHERE idRoom = ?", c.IdRoom)
	if err2 != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return false, err
	}
	for _, vals := range c.UserRooms {
		_, err4 := tx.ExecContext(ctx, `INSERT INTO USER_ROOM (idUser, idRoom) VALUES(?,?)`, vals.IdUser, vals.IdRoom)
		if err4 != nil {
			// Incase we find any error in the query execution, rollback the transaction
			tx.Rollback()
			return false, err
		}
	}
	err = nil
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return true, nil
}
