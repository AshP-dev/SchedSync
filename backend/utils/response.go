package utils

import (
    "encoding/json"
    "net/http"
)

// JSONResponse sends a JSON response with the given status code and data
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

// RespondWithError sends a JSON response with an error message
func RespondWithError(w http.ResponseWriter, status int, message string) {
    JSONResponse(w, status, map[string]string{"error": message})
}

// RespondWithJSON sends a JSON response with the given data
func RespondWithJSON(w http.ResponseWriter, status int, data interface{}) {
    JSONResponse(w, status, data)
}
