package BUSINESS

import (
	"ROOMS/MODELS"
	"ROOMS/STATICS"
	"log"
	"strings"
)

func GetBlockById(Id int) (MODELS.BLOCKS, bool) {
	var Block MODELS.BLOCKS
	db, err := STATICS.Connectdatabase()
	// Query all users
	if db == nil {

		log.Print("can not connect to database!")
		return Block, false
	}
	defer db.Close()

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

func GetBlockByIdOwner(IdOwner int) []MODELS.BLOCKS {
	var listBlock []MODELS.BLOCKS
	db, err := STATICS.Connectdatabase()
	// Query all users
	if db == nil {

		log.Print("can not connect to database!")
		return nil
	}
	defer db.Close()

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

func CreateBlock(b MODELS.BLOCKS) (bool, error) {

	db, err := STATICS.Connectdatabase()
	// Query all users
	if err != nil {

		log.Print("can not connect to database!")
		return false, err
	}
	defer db.Close()

	rows, err := db.Query(`insert into BLOCKS(nameBlock,address,description,idOwner)
							  values(?,?,?,?)`, b.NameBlock, b.Address, b.Description, b.IdOwner)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return true, nil
}

func UpdateBlock(b MODELS.BLOCKS) (bool, error) {

	db, err := STATICS.Connectdatabase()
	// Query all users
	if err != nil {

		log.Print("can not connect to database!")
		return false, err
	}
	defer db.Close()

	rows, err := db.Exec("update BLOCKS set nameBlock = ? , address = ? , description = ? where id = ?", b.NameBlock, b.Address, b.Description, b.Id)

	num, err := rows.RowsAffected()
	m := int64(num)
	if m == 0 {
		return false, err
	}
	return true, nil
}

func DeleteBlock(id int) (bool, error) {
	db, err := STATICS.Connectdatabase()

	if err != nil {
		log.Print("can not connect to database!")
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(`delete from BLOCKS where id = ?`, id)

	if err != nil {
		panic(err)
	}

	num, err := res.RowsAffected()
	m := int64(num)
	if m == 0 {
		return false, err
	}
	return true, nil
}

func DeleteBlocks(ids []int) (bool, error) {
	db, err := STATICS.Connectdatabase()

	if err != nil {
		log.Print("can not connect to database!")
		return false, err
	}
	defer db.Close()

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	stmt := `delete from BLOCKS where id in (?` + strings.Repeat(",?", len(args)-1) + `)`
	rows, err := db.Exec(stmt, args...)

	num, err := rows.RowsAffected()
	m := int64(num)
	if m == 0 {
		return false, err
	}
	return true, nil
}
