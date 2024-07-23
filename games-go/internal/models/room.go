package models

import (
	"database/sql"
)

type Room struct {
	id          int
	Name        string
	DisplayName string
	UserIds     []string
}

func CreateRoomTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS rooms (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			displayName TEXT NOT NULL

		);
	`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func CreateUserRoomsTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS user_rooms (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			room_id INTEGER NOT NULL,
			FOREIGN KEY(user_id) REFERENCES rooms(id),
			FOREIGN KEY(room_id) REFERENCES users(id)
		);
	`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
