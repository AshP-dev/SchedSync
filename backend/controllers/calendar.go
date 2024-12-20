package controllers

import (
    "encoding/json"
    "net/http"
    "ankified_planner/models"
    "ankified_planner/services"
    "ankified_planner/utils"
    "github.com/gorilla/mux"
)

func CreateCalendarEvent(w http.ResponseWriter, r *http.Request) {
    var event models.CalendarEvent
    err := json.NewDecoder(r.Body).Decode(&event)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    userID, err := services.Authenticate(r)
    if err != nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
        return
    }

    event.UserID = userID
    eventID, err := models.CreateCalendarEvent(event)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"event_id": eventID})
}

func UpdateCalendarEvent(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    eventID := vars["eventId"]

    var event models.CalendarEvent
    err := json.NewDecoder(r.Body).Decode(&event)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    userID, err := services.Authenticate(r)
    if err != nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
        return
    }

    event.UserID = userID
    updatedEvent, err := models.UpdateCalendarEvent(eventID, event)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, updatedEvent)
}

func DeleteCalendarEvent(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    eventID := vars["eventId"]

    userID, err := services.Authenticate(r)
    if err != nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
        return
    }

    err = models.DeleteCalendarEvent(eventID, userID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}