package routes

import (
	"log"
	"net/http"
	"schedsync/controllers"
	"schedsync/repositories"

	"github.com/gorilla/mux"
)

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RegisterRoutes(cardRepo repositories.CardRepository) *mux.Router {
	router := mux.NewRouter()

	// Card Management Routes
	router.HandleFunc("/api/cards", controllers.CreateCard(cardRepo)).Methods("POST")
	router.HandleFunc("/api/cards", controllers.GetCards(cardRepo)).Methods("GET")
	router.HandleFunc("/api/cards/{cardId}", controllers.UpdateCard(cardRepo)).Methods("PUT")
	router.HandleFunc("/api/cards/{cardId}", controllers.DeleteCard(cardRepo)).Methods("DELETE")

	// // Calendar Management Routes
	// router.HandleFunc("/api/calendar/events", controllers.CreateCalendarEvent(calendarEventRepo)).Methods("POST")
	// router.HandleFunc("/api/calendar/events/{eventId}", controllers.UpdateCalendarEvent(calendarEventRepo)).Methods("PUT")
	// router.HandleFunc("/api/calendar/events/{eventId}", controllers.DeleteCalendarEvent(calendarEventRepo)).Methods("DELETE")

	return router
}

func StartServer(cardRepo repositories.CardRepository) {
	router := RegisterRoutes(cardRepo)
	handler := enableCors(router)
	log.Println("Server starting on :3010...")
	if err := http.ListenAndServe(":3010", handler); err != nil {
		log.Fatal(err)
	}
}
