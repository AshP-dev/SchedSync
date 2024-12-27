package repositories

import (
	"schedsync/models"
)

type CardRepository interface {
	CreateCard(card models.Card) (string, error)
	GetCards(deckID, tags, dueDate string) ([]models.Card, error)
	UpdateCard(cardID string, card models.Card) (models.Card, error)
	DeleteCard(cardID string) error
	GetCardByID(cardID string) (models.Card, error)
	ReviewCard(cardID string, rating int) (models.Card, error)
}
