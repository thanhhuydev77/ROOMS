package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
)

func GetBlockById(Id int) (MODELS.BLOCKS, bool) {
	return DATABASE.GetBlockById(Id)
}

func GetBlockByIdOwner(db *sql.DB, IdOwner int) []MODELS.BLOCKS {
	return DATABASE.GetBlockByIdOwner(db, IdOwner)
}

func CreateBlock(b MODELS.BLOCKS) (bool, error) {
	return DATABASE.CreateBlock(b)
}

func UpdateBlock(b MODELS.BLOCKS) (bool, error) {
	return DATABASE.UpdateBlock(b)
}

func DeleteBlock(id int) (bool, error) {
	return DATABASE.DeleteBlock(id)
}

func DeleteBlocks(ids []int) (bool, error) {
	return DATABASE.DeleteBlocks(ids)
}
