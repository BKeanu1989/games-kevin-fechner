package models

import (
	"database/sql"
	"strconv"
)

type Card struct {
	Id          int
	Description string
	File        []byte
	Game        string
}

func CreateCardTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS cards (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			description TEXT NOT NULL,
			file BLOB NOT NULL,
			game_id INTEGER NOT NULL,
			FOREIGN KEY(game_id) REFERENCES game(id)
		)
	`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func AddCard(db *sql.DB, description string, file []byte) (sql.Result, error) {
	stmt, err := db.Prepare("INSERT INTO cards (description, file) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(description, file)
}

func GetCardsFromGame(db *sql.DB, game_id int) ([]Card, error) {
	stmt, err := db.Prepare("SELECT id, description, file, game_id FROM cards where game_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// res, err := db.Exec(strconv.Itoa(game_id))
	rows, err := stmt.Query(strconv.Itoa(game_id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []Card
	for rows.Next() {
		var card Card
		err := rows.Scan(&card.Id, &card.Description, &card.File, &card.Game)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}
