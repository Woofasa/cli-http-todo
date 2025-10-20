package domain

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

var ErrInvalidName = errors.New("invalid name")

type Status bool

type Task struct {
	ID          string     `json:"id" db:"id"`
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

func (t *Task) ChangeStatus(status bool) error {
	if t.Status == status {
		return fmt.Errorf("task is already %t", status)
	}
	switch t.Status {
	case true:
		t.Status = false
		currentTime := time.Now()
		t.CompletedAt = &currentTime
	case false:
		t.Status = true
		t.CompletedAt = nil
	}
	return nil
}

func (t *Task) ChangeTitle(title string) error {
	if title == "" || utf8.RuneCountInString(title) > 24 {
		return ErrInvalidName
	}
	t.Title = title
	return nil
}

func (t *Task) ChangeDescription(desc string) {
	if desc == "" {
		t.Description = "-"
	} else {
		t.Description = desc
	}
}
