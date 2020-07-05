package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
	"database/sql"
)

//get all contract with block id
func GetContractByBlockId(db *sql.DB, BlockId int) ([]MODELS.GET_CONTRACTS_REQUEST, bool, error) {
	return DATABASE.GetContractByBlockId(db, BlockId)
}

//create a new contract
func CreateContract(db *sql.DB, CCR MODELS.CREATE_UPDATE_CONTRACT_REQUEST) bool {
	return DATABASE.CreateContract(db, CCR)
}

//delete a contract
func DeleteContract(db *sql.DB, idCustomer int) (bool, error) {
	return DATABASE.DeleteContract(db, idCustomer)
}

//delete many contracts
func DeleteAllContract(db *sql.DB, idContract []int) (bool, error) {
	return DATABASE.DeleteAllContract(db, idContract)
}

//update contract
func UpdateContract(db *sql.DB, c MODELS.CREATE_UPDATE_CONTRACT_REQUEST) (bool, error) {
	return DATABASE.UpdateContract(db, c)
}
