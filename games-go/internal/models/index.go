package models

import "database/sql"

func CreateAllTables(db *sql.DB) {
	if err := CreateCardTable(db); err != nil {
		panic(err)
	}

	if err := CreateUserTable(db); err != nil {
		panic(err)
	}

	if err := CreateRoomTable(db); err != nil {
		panic(err)
	}

	if err := CreateUserTable(db); err != nil {
		panic(err)
	}

	if err := CreateUserRoomsTable(db); err != nil {
		panic(err)
	}

}
