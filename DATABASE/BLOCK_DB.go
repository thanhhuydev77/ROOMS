package DATABASE

import (
	"ROOMS/MODELS"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func GetBlockById(db *sql.DB, Id int) (MODELS.BLOCKS, bool) {
	var Block MODELS.BLOCKS
	//db, err := connectdatabase()
	//// Query all users
	if db == nil {

		log.Print("can not connect to database!")
		return Block, false
	}
	//defer db.Close()

	rows, err := db.Query("select * from BLOCKS where id = ?", Id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var block MODELS.BLOCKS
		err := rows.Scan(&block.Id, &block.NameBlock, &block.Address, &block.Description, &block.IdOwner)
		if err != nil {
			log.Fatal(err)
		}
		Block = block
	}
	defer rows.Close()
	return Block, true
}

func GetBlockByIdOwner(db *sql.DB, IdOwner int) []MODELS.BLOCKS {
	var listBlock []MODELS.BLOCKS
	//db, err := connectdatabase()
	//// Query all users
	if db == nil {

		log.Print("can not connect to database!")
		return nil
	}
	//defer db.Close()

	rows, err := db.Query("select * from BLOCKS where idowner = ?", IdOwner)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var block MODELS.BLOCKS
		err := rows.Scan(&block.Id, &block.NameBlock, &block.Address, &block.Description, &block.IdOwner)
		if err != nil {
			log.Fatal(err)
		}
		listBlock = append(listBlock, block)
	}
	defer rows.Close()
	return listBlock
}

func CreateBlock(db *sql.DB, b MODELS.BLOCKS) (bool, error) {

	//db, err := connectdatabase()
	//// Query all users
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
	rows, err := db.Query(`insert into BLOCKS(nameBlock,address,description,idOwner)
							  values(?,?,?,?)`, b.NameBlock, b.Address, b.Description, b.IdOwner)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return true, nil
}

func UpdateBlock(db *sql.DB, b MODELS.BLOCKS) (bool, error) {

	//db, err := connectdatabase()
	//// Query all users
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

	rows, err := db.Query("update BLOCKS set nameBlock = ? , address = ? , description = ? where id = ?", b.NameBlock, b.Address, b.Description, b.Id)
	//fmt.Print(err)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return true, nil
}

func DeleteBlock(db *sql.DB, id int) (bool, error) {
	//db, err := connectdatabase()
	//
	//if err != nil {
	//	log.Print("can not connect to database!")
	//	return false, err
	//}
	//defer db.Close()
	if db == nil {
		log.Print("can not connect to database!")
		return false, fmt.Errorf("can not connect db")
	}

	res, err := db.Query(`delete from BLOCKS where id = ?`, id)

	if err != nil {
		return false, err
	}
	defer res.Close()

	return true, nil
}

func DeleteBlocks(db *sql.DB, ids []int) (bool, error) {
	//db, err := connectdatabase()
	//
	//if err != nil {
	//	log.Print("can not connect to database!")
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
	stmt := `delete from BLOCKS where id in (?` + strings.Repeat(",?", len(args)-1) + `)`
	rows, err := db.Query(stmt, args...)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	return true, nil
}
