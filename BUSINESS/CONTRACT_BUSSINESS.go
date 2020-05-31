package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
)

func GetContractByBlockId(BlockId int) ([]MODELS.GET_CONTRACTS_REQUEST, bool, error) {
	return DATABASE.GetContractByBlockId(BlockId)
}

func CreateContract(CCR MODELS.CREATE_UPDATE_CONTRACT_REQUEST) bool {
	return DATABASE.CreateContract(CCR)
}

func DeleteContract(idCustomer int) (bool, error) {
	return DATABASE.DeleteContract(idCustomer)
}

func DeleteAllContract(idContract []int) (bool, error) {
	return DATABASE.DeleteAllContract(idContract)
}

func UpdateContract(c MODELS.CREATE_UPDATE_CONTRACT_REQUEST) (bool, error) {
	return DATABASE.UpdateContract(c)
}
