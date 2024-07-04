// models/user.go
package models

import (
	"database/sql"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func CreateUserTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE
    );`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func AddUser(db *sql.DB, name, email string) (sql.Result, error) {
	stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(name, email)
}

func GetUser(db *sql.DB, id int) (*User, error) {
	var user User
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := db.QueryRow(query, id)

	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
