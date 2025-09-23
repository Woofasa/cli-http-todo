package domain

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

type User struct {
	ID         string
	Name       string
	Created_at time.Time
}

func NewUser(name string) (*User, error) {
	if name == "" || utf8.RuneCountInString(name) > 14 {
		return nil, fmt.Errorf("invalid name")
	}
	return &User{
		ID:         uuid.New().String(),
		Name:       name,
		Created_at: time.Now(),
	}, nil
}
