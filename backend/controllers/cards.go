package controllers

import (
    "encoding/json"
    "net/http"
    "ankified_planner/models"
    "ankified_planner/services"
    "ankified_planner/utils"
    "github.com/gorilla/mux"
)

func CreateCard(w http.ResponseWriter, r *http.Request) {
    var card models.Card
    err := json.NewDecoder(r.Body).Decode(&card)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    userID, err := services.Authenticate(r)
    if err != nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
        return
    }

    card.UserID = userID
    cardID, err := models.CreateCard(card)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"card_id": cardID})
}

func GetCards(w http.ResponseWriter, r *http.Request) {
    userID, err := services.Authenticate(r)
    if err != nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
        return
    }

    deckID := r.URL.Query().Get("deck_id")
    tags := r.URL.Query().Get("tags")
    dueDate := r.URL.Query().Get("due_date")

    cards, err := models.GetCards(userID, deckID, tags, dueDate)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, cards)
}

func UpdateCard(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cardID := vars["cardId"]

    var card models.Card
    err := json.NewDecoder(r.Body).Decode(&card)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    userID, err := services.Authenticate(r)
    if err != nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
        return
    }

    card.UserID = userID
    updatedCard, err := models.UpdateCard(cardID, card)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, updatedCard)
}

func DeleteCard(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cardID := vars["cardId"]

    userID, err := services.Authenticate(r)
    if err != nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
        return
    }

    err = models.DeleteCard(cardID, userID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func ReviewCard(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cardID := vars["cardId"]

    var review struct {
        Rating int `json:"rating"`
    }
    err := json.NewDecoder(r.Body).Decode(&review)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    userID, err := services.Authenticate(r)
    if err != nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
        return
    }

    updatedCard, err := services.ReviewCard(cardID, userID, review.Rating)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, updatedCard)
}