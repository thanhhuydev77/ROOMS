package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
)

func GetBlockById(db *sql.DB, Id int) (MODELS.BLOCKS, bool) {
	return DATABASE.GetBlockById(db, Id)
}

func GetBlockByIdOwner(db *sql.DB, IdOwner int) []MODELS.BLOCKS {
	return DATABASE.GetBlockByIdOwner(db, IdOwner)
}

func CreateBlock(db *sql.DB, b MODELS.BLOCKS) (bool, error) {
	return DATABASE.CreateBlock(db, b)
}

func UpdateBlock(db *sql.DB, b MODELS.BLOCKS) (bool, error) {
	return DATABASE.UpdateBlock(db, b)
}

func DeleteBlock(db *sql.DB, id int) (bool, error) {
	return DATABASE.DeleteBlock(db, id)
}

func DeleteBlocks(db *sql.DB, ids []int) (bool, error) {
	return DATABASE.DeleteBlocks(db, ids)
}
