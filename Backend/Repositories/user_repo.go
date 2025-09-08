package repositories

import (
	"database/sql"
	"log"
)

func InsertUsers(db *sql.DB, username string, password string, email string, age int) error {
	query := `INSERT INTO users (username, password, email, age) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, username, password, email, age)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
/**
TO DO:
Make sure to update password in order to be hashable to protect users
**/
func CheckUserInfo(db *sql.DB, userName string, email string) (string, string, string, error) {
	var  dbUserName, dbEmail, dbPassword string

	query := `SELECT username, email, password FROM users WHERE username = $1 OR email = $2`

	err := db.QueryRow(query, userName, email).Scan(&dbUserName, &dbEmail, &dbPassword)
	if err != nil {
		return "", "", "", err
	}
	return dbUserName, dbEmail, dbPassword, err
	}