package main

import (
	"ankified_planner/routes"
	"ankified_planner/utils"
	"log"
	"net/http"
	"time"
	//"google.golang.org/api/calendar/v3"
)

func main() {
	// Initialize the database connection
	populate()
	utils.GetDB()
	//defer utils.CloseDB()

	// Serve static files from the "build" directory
	fs := http.FileServer(http.Dir("./build"))
	http.Handle("/static/", fs)

	// Register API routes
	router := routes.RegisterRoutes()
	http.Handle("/", router)

	// Start server with CORS enabled
	routes.StartServer()
}

func populate() {
	db := utils.GetDB()
	//defer utils.CloseDB()

	cards := []struct {
		Front   string
		Back    string
		DeckID  string
		Tags    string
		DueDate time.Time
	}{
		{"Chole & Rice", "Delicious Chole Masala with steamed white rice.", "Deck 1", "Tag1,Tag2", time.Now().AddDate(0, 0, 1)},
		{"Chicken Curry", "Red Gravy Chicken Curry with boneless thigh cuts.", "Deck 1", "Tag2,Tag3", time.Now().AddDate(0, 0, 2)},
		//{"Front 3", "Back 3", "Deck 2", "Tag1,Tag3", time.Now().AddDate(0, 0, 3)},
	}

	for _, card := range cards {
		query := `INSERT INTO cards (front, back, deck_id, tags, due_date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
		_, err := db.Exec(query, card.Front, card.Back, card.DeckID, card.Tags, card.DueDate, time.Now(), time.Now())
		if err != nil {
			log.Fatalf("Failed to insert card: %v", err)
		}
	}

	log.Println("Database populated with sample cards.")
}
