package main

import (
    "log"
    "net/http"
    "ankified_planner/routes"
    "ankified_planner/utils"
)

func main() {
    // Initialize the database connection
    db := utils.GetDB()
    defer utils.CloseDB()

    // Serve static files from the "build" directory
    fs := http.FileServer(http.Dir("./build"))
    http.Handle("/static/", fs)

    // Register API routes
    router := routes.RegisterRoutes()
    http.Handle("/", router)

    log.Println("Listening on :3000...")
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatal(err)
    }
}