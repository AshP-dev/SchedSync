package models

import (
	"ankified_planner/utils"
	"database/sql"
	"strconv"
	"time"
)

type CalendarEvent struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateCalendarEvent(event CalendarEvent) (string, error) {
	db := utils.GetDB()
	query := "INSERT INTO calendar_events (title, start_time, end_time, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := db.Exec(query, event.Title, event.StartTime, event.EndTime, event.UserID, time.Now(), time.Now())
	if err != nil {
		return "", err
	}
	eventID, err := result.LastInsertId()
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(eventID, 10), nil
}

func UpdateCalendarEvent(eventID string, event CalendarEvent) (CalendarEvent, error) {
	db := utils.GetDB()
	query := "UPDATE calendar_events SET title = ?, start_time = ?, end_time = ?, updated_at = ? WHERE id = ? AND user_id = ?"
	_, err := db.Exec(query, event.Title, event.StartTime, event.EndTime, time.Now(), eventID, event.UserID)
	if err != nil {
		return CalendarEvent{}, err
	}

	return GetCalendarEventByID(eventID)
}

func DeleteCalendarEvent(eventID, userID string) error {
	db := utils.GetDB()
	query := "DELETE FROM calendar_events WHERE id = ? AND user_id = ?"
	_, err := db.Exec(query, eventID, userID)
	return err
}

func GetCalendarEventByID(eventID string) (CalendarEvent, error) {
	db := utils.GetDB()
	query := "SELECT id, title, start_time, end_time, user_id, created_at, updated_at FROM calendar_events WHERE id = ?"
	var event CalendarEvent
	err := db.QueryRow(query, eventID).Scan(&event.ID, &event.Title, &event.StartTime, &event.EndTime, &event.UserID, &event.CreatedAt, &event.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return CalendarEvent{}, nil
		}
		return CalendarEvent{}, err
	}
	return event, nil
}
