package domain

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

var ErrInvalidName = errors.New("invalid name")

type Status bool

type Task struct {
	ID          string     `json:"id"  db:"id"`
	Title       string     `json:"title"  db:"title"`
	Description string     `json:"description"  db:"description"`
	Status      bool       `json:"status"  db:"status"`
	CreatedAt   time.Time  `json:"created_at"  db:"created_at"`
	CompletedAt *time.Time `json:"completed_at"  db:"completed_at"`
}

func NewTask(title string, desc string) (*Task, error) {
	if title == "" || utf8.RuneCountInString(title) > 24 {
		return nil, ErrInvalidName
	}
	if desc == "" {
		desc = "-"
	}
	return &Task{
		ID:          uuid.New().String(),
		Title:       title,
		Description: desc,
		Status:      true,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}, nil
}
