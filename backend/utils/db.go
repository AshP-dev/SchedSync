package utils

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		// Replace with your MySQL connection string
		dsn := "your-username:your-password@tcp(localhost:3306)/your-database"
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	})
	return db
}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}
}