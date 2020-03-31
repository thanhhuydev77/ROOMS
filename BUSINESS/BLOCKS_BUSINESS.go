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
		err := rows.Scan(&block.Id, &block.Name, &block.Address, &block.Description, &block.IdOwner)
		if err != nil {
			log.Fatal(err)
		}
		listBlock = append(listBlock, block)
	}
	defer rows.Close()
	return listBlock
}
