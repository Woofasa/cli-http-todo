package domain

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

type User struct {
	ID         string    `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"password"  db:"password"`
	Created_at time.Time `json:"created_at" db:"created_at"`
}

func NewUser(name, email, password string) (*User, error) {
	if name == "" || utf8.RuneCountInString(name) > 14 {
		return nil, fmt.Errorf("invalid name")
	}
	if email == "" {
		return nil, fmt.Errorf("invalid email")
	}
	if utf8.RuneCountInString(password) < 6 {
		return nil, fmt.Errorf("invalid password")
	}
	return &User{
		ID:         uuid.New().String(),
		Name:       name,
		Email:      email,
		Password:   password,
		Created_at: time.Now(),
	}, nil
}
