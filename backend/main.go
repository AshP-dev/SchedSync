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
	db := utils.GetDB()
	//defer utils.CloseDB()

	// Create repository instances
	cardRepo := repositories.NewSQLiteCardRepository(db)
	populateSQLite(cardRepo)
	//calendarEventRepo := repositories.NewSQLiteCalendarEventRepository(db)

	// Register API routes with repositories
	router := routes.RegisterRoutes(cardRepo)
	http.Handle("/", router)

	// Start server with CORS enabled
	routes.StartServer(cardRepo)
}

// func main() {
// 	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
// 	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
// 	opts := options.Client().ApplyURI("mongodb+srv://<SCHEDSYNC_ADMIN>:<SCHEDSYNC_ADMIN_PASSWORD>@cluster0.gdzbr.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)

// 	// Create a new client and connect to the server
// 	client, err := mongo.Connect(context.TODO(), opts)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer func() {
// 		if err = client.Disconnect(context.TODO()); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	// Send a ping to confirm a successful connection
// 	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

// 	// Create repository instances
// 	cardCollection := client.Database("schedsync").Collection("cards")
// 	cardRepo := repositories.NewMongoCardRepository(cardCollection)
// 	populateMongo(cardCollection)

// 	// Register API routes with repositories
// 	router := routes.RegisterRoutes(cardRepo)
// 	http.Handle("/", router)
// 	c := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"http://localhost:3000"}, // Update with your frontend URL
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
// 		AllowedHeaders:   []string{"Content-Type"},
// 		AllowCredentials: true,
// 	})

// 	// Start server with CORS enabled
// 	handler := c.Handler(router)
// 	log.Fatal(http.ListenAndServe(":8080", handler))
// 	// Start server with CORS enabled
// 	// routes.StartServer(cardRepo)
// }

func populateSQLite(repo repositories.CardRepository) {
	cards := []models.Card{
		{Front: "Chole & Rice", Back: "Delicious Chole Masala with steamed white rice.", DeckID: "Deck 1", Tags: "Tag1,Tag2", DueDate: time.Now().AddDate(0, 0, 1)},
		{Front: "Chicken Curry", Back: "Red Gravy Chicken Curry with boneless thigh cuts.", DeckID: "Deck 1", Tags: "Tag2,Tag3", DueDate: time.Now().AddDate(0, 0, 2)},
		{Front: "Pasta", Back: "Red Sauce Pasta with grilled chicken", DeckID: "Deck 2", Tags: "Tag1,Tag3", DueDate: time.Now().AddDate(0, 0, 3)},
	}

	for _, card := range cards {
		_, err := repo.CreateCard(card)
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
