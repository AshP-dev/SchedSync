package repositories

import (
	"database/sql"
	"schedsync/models"
	"strconv"
	"time"
)

type SQLiteCardRepository struct {
	db *sql.DB
}

func NewSQLiteCardRepository(db *sql.DB) *SQLiteCardRepository {
	return &SQLiteCardRepository{db: db}
}

func (r *SQLiteCardRepository) CreateCard(card models.Card) (string, error) {
	query := "INSERT INTO cards (front, back, deck_id, tags, due_date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, card.Front, card.Back, card.DeckID, card.Tags, card.DueDate, time.Now(), time.Now())
	if err != nil {
		return "", err
	}
	cardID, err := result.LastInsertId()
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(cardID, 10), nil
}

func (r *SQLiteCardRepository) DeleteCard(id string) error {
	query := `DELETE FROM cards WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *SQLiteCardRepository) GetCardByID(cardID string) (models.Card, error) {
	var card models.Card
	query := "SELECT id, front, back, deck_id, tags, due_date, created_at, updated_at FROM cards WHERE id = ?"
	err := r.db.QueryRow(query, cardID).Scan(&card.ID, &card.Front, &card.Back, &card.DeckID, &card.Tags, &card.DueDate, &card.CreatedAt, &card.UpdatedAt)
	if err != nil {
		return models.Card{}, err
	}
	return card, nil
}

func (r *SQLiteCardRepository) GetCards(deckID string, tag string, dueDate string) ([]models.Card, error) {
	query := "SELECT id, front, back, deck_id, tags, due_date, created_at, updated_at FROM cards WHERE 1=1"
	args := []interface{}{}

	if deckID != "" {
		query += " AND deck_id = ?"
		args = append(args, deckID)
	}
	if tag != "" {
		query += " AND tags LIKE ?"
		args = append(args, "%"+tag+"%")
	}
	if dueDate != "" {
		query += " AND due_date = ?"
		args = append(args, dueDate)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []models.Card
	for rows.Next() {
		var card models.Card
		if err := rows.Scan(&card.ID, &card.Front, &card.Back, &card.DeckID, &card.Tags, &card.DueDate, &card.CreatedAt, &card.UpdatedAt); err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (r *SQLiteCardRepository) ReviewCard(cardID string, rating int) (models.Card, error) {
	query := "UPDATE cards SET reviewed_at = ? WHERE id = ?"
	_, err := r.db.Exec(query, time.Now(), cardID)
	if err != nil {
		return models.Card{}, err
	}

	return r.GetCardByID(cardID)
}

func (r *SQLiteCardRepository) UpdateCard(cardID string, card models.Card) (models.Card, error) {
	query := "UPDATE cards SET front = ?, back = ?, deck_id = ?, tags = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.Exec(query, card.Front, card.Back, card.DeckID, card.Tags, time.Now(), cardID)
	if err != nil {
		return models.Card{}, err
	}

	return r.GetCardByID(cardID)
}

// type SQLiteCalendarEventRepository struct {
// 	db *sql.DB
// }

// func NewSQLiteCalendarEventRepository(db *sql.DB) *SQLiteCalendarEventRepository {
// 	return &SQLiteCalendarEventRepository{db: db}
// }

// func (r *SQLiteCalendarEventRepository) CreateCalendarEvent(event models.CalendarEvent) (string, error) {
// 	query := "INSERT INTO calendar_events (title, start_time, end_time, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
// 	result, err := r.db.Exec(query, event.Title, event.StartTime, event.EndTime, event.UserID, time.Now(), time.Now())
// 	if err != nil {
// 		return "", err
// 	}
// 	eventID, err := result.LastInsertId()
// 	if err != nil {
// 		return "", err
// 	}
// 	return strconv.FormatInt(eventID, 10), nil
// }
