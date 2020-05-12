package BUSINESS

import (
	"ROOMS/MODELS"
	"ROOMS/STATICS"
	"log"
)

func GetCustomersByUserId(userId int)([]MODELS.CUSTOMER,bool, error)  {
	var listCustomers []MODELS.CUSTOMER
	db, err := STATICS.Connectdatabase()

	if err != nil{
		return nil,false, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT CU.*, R.nameRoom FROM CUSTOMERS CU LEFT JOIN USER_ROOM UR ON CU.id = UR.idUser " +
		"LEFT JOIN ROOMS R ON UR.idRoom = R.id  WHERE idOwner = ?", userId)
	if err != nil{
		log.Fatal(err)
		return nil,false, err
	}

	for rows.Next(){
		var c MODELS.CUSTOMER

		err := rows.Scan(&c.Id, &c.CodeUser, &c.UserName, &c.Pass, &c.FullName,
			&c.IdentifyFront, &c.IdentifyBack, &c.DateBirth, &c.Address,
			&c.Role, &c.Sex, &c.Job, &c.WorkPlace, &c.TempReg, &c.Province,
			&c.Email, &c.Avatar, &c.PhoneNumber, &c.IdOwner, &c.Note, &c.NameRoom)

		if err != nil{
			log.Fatal(err)
			return nil,false, err
		}
		listCustomers = append(listCustomers, c)
	}
	defer rows.Close()
	return listCustomers,true,nil
}
