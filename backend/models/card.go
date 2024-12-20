package models

import (
    "database/sql"
    "time"
    "ankified_planner/utils"
)

type Card struct {
    ID        string    `json:"id"`
    Front     string    `json:"front"`
    Back      string    `json:"back"`
    DeckID    string    `json:"deck_id"`
    Tags      string    `json:"tags"`
    UserID    string    `json:"user_id"`
    DueDate   time.Time `json:"due_date"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func CreateCard(card Card) (string, error) {
    db := utils.GetDB()
    query := "INSERT INTO cards (front, back, deck_id, tags, user_id, due_date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
    result, err := db.Exec(query, card.Front, card.Back, card.DeckID, card.Tags, card.UserID, card.DueDate, time.Now(), time.Now())
    if err != nil {
        return "", err
    }
    cardID, err := result.LastInsertId()
    if err != nil {
        return "", err
    }
    return string(cardID), nil
}

func GetCards(userID, deckID, tags, dueDate string) ([]Card, error) {
    db := utils.GetDB()
    query := "SELECT id, front, back, deck_id, tags, user_id, due_date, created_at, updated_at FROM cards WHERE user_id = ?"
    args := []interface{}{userID}

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
        err := rows.Scan(&card.ID, &card.Front, &card.Back, &card.DeckID, &card.Tags, &card.UserID, &card.DueDate, &card.CreatedAt, &card.UpdatedAt)
        if err != nil {
            return nil, err
        }
        cards = append(cards, card)
    }
    return cards, nil
}

func UpdateCard(cardID string, card Card) (Card, error) {
    db := utils.GetDB()
    query := "UPDATE cards SET front = ?, back = ?, deck_id = ?, tags = ?, updated_at = ? WHERE id = ? AND user_id = ?"
    _, err := db.Exec(query, card.Front, card.Back, card.DeckID, card.Tags, time.Now(), cardID, card.UserID)
    if err != nil {
        return Card{}, err
    }

    return GetCardByID(cardID)
}

func DeleteCard(cardID, userID string) error {
    db := utils.GetDB()
    query := "DELETE FROM cards WHERE id = ? AND user_id = ?"
    _, err := db.Exec(query, cardID, userID)
    return err
}

func GetCardByID(cardID string) (Card, error) {
    db := utils.GetDB()
    query := "SELECT id, front, back, deck_id, tags, user_id, due_date, created_at, updated_at FROM cards WHERE id = ?"
    var card Card
    err := db.QueryRow(query, cardID).Scan(&card.ID, &card.Front, &card.Back, &card.DeckID, &card.Tags, &card.UserID, &card.DueDate, &card.CreatedAt, &card.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return Card{}, nil
        }
        return Card{}, err
    }
    return card, nil
}

func ReviewCard(cardID, userID string, rating int) (Card, error) {
    db := utils.GetDB()
    card, err := GetCardByID(cardID)
    if err != nil {
        return Card{}, err
    }

    // Update the due date based on the rating using a spaced repetition algorithm
    newDueDate := utils.CalculateNewDueDate(card.DueDate, rating)
    query := "UPDATE cards SET due_date = ?, updated_at = ? WHERE id = ? AND user_id = ?"
    _, err = db.Exec(query, newDueDate, time.Now(), cardID, userID)
    if err != nil {
        return Card{}, err
    }

    return GetCardByID(cardID)
}
