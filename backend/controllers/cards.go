package controllers

import (
	"ankified_planner/models"
	"ankified_planner/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateCard(w http.ResponseWriter, r *http.Request) {
	var card models.Card
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	cardID, err := models.CreateCard(card)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"card_id": cardID})
}

func GetCards(w http.ResponseWriter, r *http.Request) {
	deckID := r.URL.Query().Get("deck_id")
	tags := r.URL.Query().Get("tags")
	dueDate := r.URL.Query().Get("due_date")

	cards, err := models.GetCards(deckID, tags, dueDate)
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

	err := models.DeleteCard(cardID)
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

	updatedCard, err := models.ReviewCard(cardID, review.Rating)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, updatedCard)
}
