package repositories

import (
	"database/sql"
	"log"
)

func InsertFriend(db *sql.DB, userID, friendID int) error {
	query := `INSERT INTO friends (user_ID, friend_ID) VALUES ($1, $2)`
	_, err := db.Exec(query, userID, friendID)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
