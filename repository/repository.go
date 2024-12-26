package repository

import (
	"database/sql"
	"goauth/models"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if (err != nil) {
		log.Fatal(err)
	}

	// Create users table if it doesn't exist
	createTableQuery := `CREATE TABLE IF NOT EXISTS users (
		uuid TEXT PRIMARY KEY,
		username TEXT UNIQUE,
		password TEXT
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT uuid, username, password FROM users WHERE username = ?", username).Scan(&user.UUID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	_, err := db.Exec("INSERT INTO users (uuid, username, password) VALUES (?, ?, ?)", user.UUID, user.Username, user.Password)
	return err
}
