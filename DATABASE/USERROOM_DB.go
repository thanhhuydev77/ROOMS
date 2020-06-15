package DATABASE

import "strings"

func DeleteManyUserRoom(ids []int) (bool, error)  {
	db, err := connectdatabase()
	if err != nil {
		return false, err
	}
	defer db.Close()

	args := make([]interface{} , len(ids))
	for i, id := range ids {
		args[i] = id
	}

	stmt := `DELETE FROM USER_ROOM id in (?` + strings.Repeat(",?", len(args)-1) + `)`
	rows, err := db.Exec(stmt, args...)

	num, err := rows.RowsAffected()
	m := int64(num)

	if m > 0 {
		return true, nil
	} else {
		return false, err
	}
}