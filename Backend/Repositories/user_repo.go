package repositories

import (
	"database/sql"
	"log"
)

func insertUser(db *sql.DB, username string, password string, email string, age int) error {
	query := `INSERT INTO users (username, password, email, age) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, username, password, email, age)
	if err != nil {
		log.Fatal(err)
	}
	return err
}