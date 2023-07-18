package db

import (
	_ "github.com/go-sql-driver/mysql"
)

func InsertUser(username string, password string) error {
	db := GetDatabase()
	_, err := db.Exec(`INSERT INTO user (username,password) VALUES(?,?);`, username, password)
	return err
}

func GetUsername(username string) (string, error) {
	database := GetDatabase()
	var name string
	err := database.QueryRow("SELECT username FROM user WHERE username = ?", username).Scan(&name)
	return name, err
}

func GetPasswordFromUsername(username string) (string, error) {
	database := GetDatabase()
	var password string
	err := database.QueryRow("SELECT password FROM user WHERE username= ?", username).Scan(&password)
	return password, err
}

func GetUserGuild(username string) (string, error) {
	database := GetDatabase()
	var guid string
	err := database.QueryRow("SELECT guid FROM user WHERE username= ?", username).Scan(&guid)
	return guid, err
}
