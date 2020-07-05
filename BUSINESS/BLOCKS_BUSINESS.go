package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
)

//get block by its id
func GetBlockById(db *sql.DB, Id int) (MODELS.BLOCKS, bool) {
	return DATABASE.GetBlockById(db, Id)
}

//get all block of user by id
func GetBlockByIdOwner(db *sql.DB, IdOwner int) []MODELS.BLOCKS {
	return DATABASE.GetBlockByIdOwner(db, IdOwner)
}

//create a new block
func CreateBlock(db *sql.DB, b MODELS.BLOCKS) (bool, error) {
	return DATABASE.CreateBlock(db, b)
}

//update a block
func UpdateBlock(db *sql.DB, b MODELS.BLOCKS) (bool, error) {
	return DATABASE.UpdateBlock(db, b)
}

//delete a block
func DeleteBlock(db *sql.DB, id int) (bool, error) {
	return DATABASE.DeleteBlock(db, id)
}

//delete many block
func DeleteBlocks(db *sql.DB, ids []int) (bool, error) {
	return DATABASE.DeleteBlocks(db, ids)
}
