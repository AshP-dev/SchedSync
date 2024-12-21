package services

import (
	"time"
)

// Task represents a task with a due date and repetition interval
type Task struct {
	ID       string
	Name     string
	DueDate  time.Time
	Interval int // Interval in days
}

// CalculateNextDueDate calculates the next due date for a task based on spaced repetition
func CalculateNextDueDate(task Task) time.Time {
	return task.DueDate.AddDate(0, 0, task.Interval)
}

// UpdateTaskInterval updates the interval for a task based on user feedback
func UpdateTaskInterval(task *Task, success bool) {
	if success {
		task.Interval *= 2 // Double the interval if the task was successful
	} else {
		task.Interval = 1 // Reset the interval if the task was not successful
	}
}

// CalculateNewDueDate calculates the new due date based on the current due date and interval
func CalculateNewDueDate(currentDueDate time.Time, interval int) time.Time {
	return currentDueDate.AddDate(0, 0, interval)
}
