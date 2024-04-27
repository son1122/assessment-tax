package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func SetDB(database *sql.DB) {
	DB = database
}
func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}

// CloseDB closes the database connection.
func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Println("Failed to close database connection:", err)
		} else {
			log.Println("Database connection closed successfully.")
		}
	}
}
