package BUSINESS

import (
	"ROOMS/MODELS"
	"ROOMS/STATICS"
	"log"
)

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

	rows, err := db.Exec(`update BLOCKS
								  set nameBlock = ? and address = ? and description = ?
								  where id = ?`, b.NameBlock, b.Address, b.Description, b.Id)

	num, err := rows.RowsAffected()
	m := int64(num)
	if m == 0 {
		return false, err
	}
	return true, nil
}

func DeleteBlock(id int) (bool, error)  {
	db, err := STATICS.Connectdatabase()

	if err != nil{
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