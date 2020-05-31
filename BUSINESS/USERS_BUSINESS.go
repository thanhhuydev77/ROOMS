package BUSINESS

import (
	"ROOMS/DATABASE"
	"ROOMS/MODELS"
)

func Login(username string, pass string) (bool, bool, MODELS.USERS) {
	return DATABASE.Login(username, pass)
}

func Register(user MODELS.RequestRegister) (bool, error) {
	return DATABASE.Register(user)
}

func GetAllUserName() []string {
	return DATABASE.GetAllUserName()
}

func GetUsers(Id int) []MODELS.USERS {
	return DATABASE.GetUsers(Id)
}
