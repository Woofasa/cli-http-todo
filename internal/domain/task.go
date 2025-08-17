package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrInvalidName = errors.New("invalid name")
var ErrAlreadyClosed = errors.New("task is already closed")
var ErrAlreadyOpened = errors.New("task is already opened")

type Status bool

var (
	Opened Status = true
	Closed Status = false
)

type Task struct {
	ID          string
	Title       string
	Description string
	Status      Status
	CreatedAt   time.Time
	CompletedAt time.Time
}

func NewTask(title string, desc string) (*Task, error) {
	if title == "" || len(title) > 20 {
		return nil, ErrInvalidName
	}
	return &Task{
		ID:          uuid.New().String(),
		Title:       title,
		Description: desc,
		Status:      Opened,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}, nil
}

func (t Task) GetID() string {
	return t.ID
}

func (t *Task) Rename(newTitle string) error {
	if newTitle == "" {
		return ErrInvalidName
	}
	t.Title = newTitle
	return nil
}

func (t *Task) ChangeDescription(newDesc string) {
	if newDesc == "" {
		t.Description = "empty"
		return
	}
	t.Description = newDesc
}

func (t *Task) CloseTask() error {
	if t.Status == Closed {
		return ErrAlreadyClosed
	}
	t.Status = Closed
	return nil
}

func (t *Task) OpenTask() error {
	if t.Status == Opened {
		return ErrAlreadyOpened
	}
	t.Status = Opened
	return nil
}
