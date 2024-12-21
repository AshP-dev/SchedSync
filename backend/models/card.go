package models

import (
	"ankified_planner/services"
	"ankified_planner/utils"
	"database/sql"
	"strconv"
	"time"
)

type Card struct {
	ID        string    `json:"id"`
	Front     string    `json:"front"`
	Back      string    `json:"back"`
	DeckID    string    `json:"deck_id"`
	Tags      string    `json:"tags"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Interval  int
}

func CreateCard(card Card) (string, error) {
	db := utils.GetDB()
	query := "INSERT INTO cards (front, back, deck_id, tags, due_date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := db.Exec(query, card.Front, card.Back, card.DeckID, card.Tags, card.DueDate, time.Now(), time.Now())
	if err != nil {
		return "", err
	}
	cardID, err := result.LastInsertId()
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(cardID, 10), nil
}

func GetCards(deckID, tags, dueDate string) ([]Card, error) {
	db := utils.GetDB()
	query := "SELECT id, front, back, deck_id, tags, due_date, created_at, updated_at FROM cards WHERE 1=1"
	args := []interface{}{}

	if deckID != "" {
		query += " AND deck_id = ?"
		args = append(args, deckID)
	}
	if tags != "" {
		query += " AND tags LIKE ?"
		args = append(args, "%"+tags+"%")
	}
	if dueDate != "" {
		query += " AND due_date <= ?"
		args = append(args, dueDate)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []Card
	for rows.Next() {
		var card Card
		err := rows.Scan(&card.ID, &card.Front, &card.Back, &card.DeckID, &card.Tags, &card.DueDate, &card.CreatedAt, &card.UpdatedAt)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func UpdateCard(cardID string, card Card) (Card, error) {
	db := utils.GetDB()
	query := "UPDATE cards SET front = ?, back = ?, deck_id = ?, tags = ?, updated_at = ? WHERE id = ?"
	_, err := db.Exec(query, card.Front, card.Back, card.DeckID, card.Tags, time.Now(), cardID)
	if err != nil {
		return Card{}, err
	}

	return GetCardByID(cardID)
}

func DeleteCard(cardID string) error {
	db := utils.GetDB()
	query := "DELETE FROM cards WHERE id = ?"
	_, err := db.Exec(query, cardID)
	return err
}

func GetCardByID(cardID string) (Card, error) {
	db := utils.GetDB()
	query := "SELECT id, front, back, deck_id, tags, due_date, created_at, updated_at FROM cards WHERE id = ?"
	var card Card
	err := db.QueryRow(query, cardID).Scan(&card.ID, &card.Front, &card.Back, &card.DeckID, &card.Tags, &card.DueDate, &card.CreatedAt, &card.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return Card{}, nil
		}
		return Card{}, err
	}
	return card, nil
}

func ReviewCard(cardID string, rating int) (Card, error) {
	db := utils.GetDB()
	card, err := GetCardByID(cardID)
	if err != nil {
		return Card{}, err
	}

	// Update the due date based on the rating using a spaced repetition algorithm
	newDueDate := services.CalculateNewDueDate(card.DueDate, rating)
	query := "UPDATE cards SET due_date = ?, updated_at = ? WHERE id = ?"
	_, err = db.Exec(query, newDueDate, time.Now(), cardID)
	if err != nil {
		return Card{}, err
	}

	return GetCardByID(cardID)
}

// UpdateDueDate updates the due date of the card based on spaced repetition
func (c *Card) UpdateDueDate(success bool) {
	if success {
		c.Interval *= 2
	} else {
		c.Interval = 1
	}
	c.DueDate = services.CalculateNewDueDate(c.DueDate, c.Interval)
}
