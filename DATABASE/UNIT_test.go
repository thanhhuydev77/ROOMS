package DATABASE

//func TestUnit(t *testing.T) {
//	//db, mock, err := sqlmock.New()
//	//if err != nil {
//	//	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	//}
//	//defer db.Close()
//	//
//	//
//	//rows := sqlmock.NewRows([]string{"id", "name", "description"}).
//	//	AddRow(1, "post 1", "hello").
//	//	AddRow(2, "post 2", "world")
//	//
//	//mock.ExpectQuery(`select \* from UNITS`).WillReturnRows(rows)
//	//
//	//
//	//
//	//// now we execute our method
//	//if result,ok := GetAllUnits(db); !ok ||len(result) == 0 {
//	//	t.Errorf("error was not expected while updating stats: %s", err)
//	//}
//	//
//	//// we make sure that all expectations were met
//	//if err := mock.ExpectationsWereMet(); err != nil {
//	//	t.Errorf("there were unfulfilled expectations: %s", err)
//	//}
//}
