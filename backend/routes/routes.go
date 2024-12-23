package routes

import (
	"ankified_planner/controllers"
	"log"
	"net/http"

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

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	// Card Management Routes
	router.HandleFunc("/api/cards", controllers.CreateCard).Methods("POST")
	router.HandleFunc("/api/cards", controllers.GetCards).Methods("GET")
	router.HandleFunc("/api/cards/{cardId}", controllers.UpdateCard).Methods("PUT")
	router.HandleFunc("/api/cards/{cardId}", controllers.DeleteCard).Methods("DELETE")

	// Calendar Management Routes
	router.HandleFunc("/api/calendar/events", controllers.CreateCalendarEvent).Methods("POST")
	router.HandleFunc("/api/calendar/events/{eventId}", controllers.UpdateCalendarEvent).Methods("PUT")
	router.HandleFunc("/api/calendar/events/{eventId}", controllers.DeleteCalendarEvent).Methods("DELETE")

	return router
}

func StartServer() {
	router := RegisterRoutes()
	handler := enableCors(router)
	log.Println("Server starting on :3010...")
	if err := http.ListenAndServe(":3010", handler); err != nil {
		log.Fatal(err)
	}
}
