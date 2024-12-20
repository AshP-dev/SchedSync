package routes

import (
    "ankified_planner/controllers"
    "github.com/gorilla/mux"
    "net/http"
)

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
    http.Handle("/", router)
    http.ListenAndServe(":3000", nil)
}