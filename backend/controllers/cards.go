package controllers

import (
	"encoding/json"
	"net/http"
	"schedsync/models"
	"schedsync/repositories"
	"schedsync/utils"

	"github.com/gorilla/mux"
)

func CreateCard(repo repositories.CardRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var card models.Card
		err := json.NewDecoder(r.Body).Decode(&card)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		cardID, err := repo.CreateCard(card)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"card_id": cardID})
	}
}

func GetCards(repo repositories.CardRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deckID := r.URL.Query().Get("deck_id")
		tags := r.URL.Query().Get("tags")
		dueDate := r.URL.Query().Get("due_date")

		cards, err := repo.GetCards(deckID, tags, dueDate)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, cards)
	}
}

func UpdateCard(repo repositories.CardRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cardID := vars["cardId"]

		var card models.Card
		err := json.NewDecoder(r.Body).Decode(&card)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		updatedCard, err := repo.UpdateCard(cardID, card)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, updatedCard)
	}
}

func DeleteCard(repo repositories.CardRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cardID := vars["cardId"]

		err := repo.DeleteCard(cardID)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	}
}

func ReviewCard(repo repositories.CardRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		updatedCard, err := repo.ReviewCard(cardID, review.Rating)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, updatedCard)
	}
}
