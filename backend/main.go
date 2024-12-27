package main

import (
	"context"
	"log"
	"net/http"
	"schedsync/models"
	"schedsync/repositories"
	"schedsync/routes"
	"schedsync/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// Initialize the database connection
	populateSQLite()
	db := utils.GetDB()
	defer utils.CloseDB()

	// Create repository instances
	cardRepo := repositories.NewSQLiteCardRepository(db)
	//calendarEventRepo := repositories.NewSQLiteCalendarEventRepository(db)

	// Register API routes with repositories
	router := routes.RegisterRoutes(cardRepo)
	http.Handle("/", router)

	// Start server with CORS enabled
	routes.StartServer(cardRepo)
}

func populateSQLite() {
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
		{"Pasta", "Red Sauce Pasta with grilled chicken", "Deck 2", "Tag1,Tag3", time.Now().AddDate(0, 0, 3)},
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

func populateMongo(collection *mongo.Collection) {
	cards := []interface{}{
		models.Card{
			Front:     "Chole & Rice",
			Back:      "Delicious Chole Masala with steamed white rice.",
			DeckID:    "Deck 1",
			Tags:      "Tag1,Tag2",
			DueDate:   time.Now().AddDate(0, 0, 1),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		models.Card{
			Front:     "Chicken Curry",
			Back:      "Red Gravy Chicken Curry with boneless thigh cuts.",
			DeckID:    "Deck 1",
			Tags:      "Tag2,Tag3",
			DueDate:   time.Now().AddDate(0, 0, 2),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		models.Card{
			Front:     "Pasta",
			Back:      "Red Sauce Pasta with grilled chicken",
			DeckID:    "Deck 2",
			Tags:      "Tag1,Tag3",
			DueDate:   time.Now().AddDate(0, 0, 3),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	_, err := collection.InsertMany(context.Background(), cards)
	if err != nil {
		log.Fatalf("Failed to insert cards: %v", err)
	}

	log.Println("Database populated with sample cards.")
}
