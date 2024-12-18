package main

import (
    "context" //context libraries for API calls
    "database/sql" // DB
    "encoding/json" //Json handling
    "log" // logger
    "net/http" // http client and server implementation
    "os" // os functions

    "github.com/go-sql-driver/mysql" // mysql driver for go
    // "github.com/grok-ai/grok-go-client" 
    "google.golang.org/api/calendar/v3" // google calendar API
    "google.golang.org/api/option" //options for api clients
)

var db *sql.DB
var grokClient *grok.Client

func main() {
    // Connect to MySQL database
    db = initDB()
    defer db.Close()

    // Initialize Grok LLM API client
    grokClient = grok.NewClient("your-grok-api-key")

    // Serve static files from the "build" directory
    fs := http.FileServer(http.Dir("./build"))
    http.Handle("/static/", fs)

    // API endpoints
    http.HandleFunc("/api/schedule", handleSchedule)
    http.HandleFunc("/api/approve", handleApproval)

    // Handle all other routes by serving the index.html file
    http.HandleFunc("/", serveIndex)

    log.Println("Listening on :3000...")
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatal(err)
    }
}

// Initialising the Database

func initDB() *sql.DB {
    cfg := mysql.Config{
        User:                 "your-username",
        Passwd:               "your-password",
        Net:                  "tcp",
        Addr:                 "localhost:3306",
        DBName:               "your-database",
        AllowNativePasswords: true,
    }

    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }

    // Verify connection
    err = db.Ping()
    if err != nil {
        log.Fatal("Database ping failed:", err)
    }

    log.Println("Connected to the database")
    return db
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
    path := "./build" + r.URL.Path
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        http.ServeFile(w, r, "./build/index.html")
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    } else {
        http.ServeFile(w, r, path)
    }
}

func handleSchedule(w http.ResponseWriter, r *http.Request) {
    // Retrieve Google Calendar events
    ctx := context.Background()
    calendarService, err := calendar.NewService(ctx, option.WithAPIKey("your-google-api-key"))
    if err != nil {
        http.Error(w, "Failed to create Calendar service", http.StatusInternalServerError)
        return
    }

    events, err := calendarService.Events.List("primary").Do()
    if err != nil {
        http.Error(w, "Failed to retrieve events", http.StatusInternalServerError)
        return
    }

    // Collect user input
    var userInput struct {
        Preferences string `json:"preferences"`
    }
    err = json.NewDecoder(r.Body).Decode(&userInput)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Generate schedule suggestions using Grok LLM API
    suggestions, err := grokClient.GenerateSchedule(events, userInput.Preferences)
    if err != nil {
        http.Error(w, "Failed to generate schedule", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(suggestions)
}

func handleApproval(w http.ResponseWriter, r *http.Request) {
    // Parse approval data
    var approval struct {
        TaskID   string `json:"taskId"`
        Approved bool   `json:"approved"`
    }
    err := json.NewDecoder(r.Body).Decode(&approval)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Store approval in the database
    _, err = db.Exec("INSERT INTO approvals (task_id, approved) VALUES (?, ?)", approval.TaskID, approval.Approved)
    if err != nil {
        http.Error(w, "Failed to save approval", http.StatusInternalServerError)
        return
    }

    // Update Grok LLM cache
    err = grokClient.UpdateCache(approval.TaskID, approval.Approved)
    if err != nil {
        http.Error(w, "Failed to update Grok cache", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}