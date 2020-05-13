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

func CreateCustomer(c MODELS.CUSTOMER_INPUT)(bool, error)  {
	db, err := STATICS.Connectdatabase()
	if err != nil{
		log.Fatalln(err)
		return false, err
	}
	defer db.Close()

	rs, err := db.Query(`INSERT INTO CUSTOMERS (codeUser, fullName, identifyFront, identifyBack, dateBirth, sex, 
		job, workPlace, tempReg, email, avatar, phoneNumber, idOwner, note) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)` ,
		c.CodeUser, c.FullName, c.IdentifyFront, c.IdentifyBack, c.DateBirth, c.Sex, c.Job,
		c.WorkPlace, c.TempReg, c.Email, c.Avatar, c.PhoneNumber, c.IdOwner, c.Note)


	if err != nil{
		log.Fatalln(err)
		return false, err
	}
	defer rs.Close()

	return true, nil
}

func DeleteCustomer(idCustomer int)(bool, error)  {
	db, err := STATICS.Connectdatabase()
	if err != nil{
		log.Fatalln(err)
		return false, err
	}
	defer db.Close()

	rs, err := db.Query(`DELETE FROM CUSTOMERS WHERE id = ?`, idCustomer)

	if err != nil{
		log.Fatalln(err)
		return false, err
	}
	defer rs.Close()

	return true, nil
}