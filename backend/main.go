package main

import (
	"ankified_planner/routes"
	"ankified_planner/utils"
	"log"
	"net/http"
	//"google.golang.org/api/calendar/v3"
)

func main() {
	// Initialize the database connection
	//db := utils.GetDB()
	defer utils.CloseDB()

	// Serve static files from the "build" directory
	fs := http.FileServer(http.Dir("./build"))
	http.Handle("/static/", fs)

	// Register API routes
	router := routes.RegisterRoutes()
	http.Handle("/", router)

	log.Println("Listening on :3010...")
	err := http.ListenAndServe(":3010", nil)
	if err != nil {
		log.Fatal(err)
	}
}
