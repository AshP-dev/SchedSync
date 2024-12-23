package utils

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db   *sql.DB
	once sync.Once
)

type Card struct {
	ID        string    `json:"id"`
	Front     string    `json:"front"`
	Back      string    `json:"back"`
	DeckID    string    `json:"deck_id"`
	Tags      string    `json:"tags"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Interval  int
	ImageURL  string `json:"image_url"`
	LinkURL   string `json:"link_url"`
}

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		// SQLite connection string
		dsn := "file:cards.db?cache=shared&mode=rwc"
		db, err = sql.Open("sqlite3", dsn)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Create tables if they don't exist
		createTables()
	})
	return db
}

func createTables() {
	createCardsTable := `
	CREATE TABLE IF NOT EXISTS cards (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		front TEXT,
		back TEXT,
		deck_id TEXT,
		tags TEXT,
		due_date DATETIME,
		created_at DATETIME,
		updated_at DATETIME
	);`
	_, err := db.Exec(createCardsTable)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}
}

func GetCards(w http.ResponseWriter, r *http.Request) {
	db := GetDB()
	rows, err := db.Query("SELECT id, front, back, deck_id, tags, due_date, image_url, link_url FROM cards")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cards []Card
	for rows.Next() {
		var card Card
		if err := rows.Scan(&card.ID, &card.Front, &card.Back, &card.DeckID, &card.Tags, &card.DueDate, &card.ImageURL, &card.LinkURL); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cards = append(cards, card)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cards); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
