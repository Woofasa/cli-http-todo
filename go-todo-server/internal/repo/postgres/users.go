package postgres

import (
	"context"
	"errors"
	"fmt"
	"main/internal/domain"

	_ "github.com/lib/pq"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type Users struct {
	db *DB
}

func NewUserStorage(db *DB) *Users {
	return &Users{
		db: db,
	}
}

func (s *Users) Init(ctx context.Context) error {
	q := `CREATE TABLE IF NOT EXISTS users(
		id UUID PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL
	);`

	_, err := s.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create user table: %w", err)
	}

	return nil
}

func (s *Users) GetUsers(ctx context.Context) ([]*domain.User, error) {
	q := `SELECT id, name, email, password, created_at FROM users ORDER BY created_at;`

	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("postgres user query: %w", err)
	}
	defer rows.Close()

	users := make([]*domain.User, 0)
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Created_at); err != nil {
			return nil, fmt.Errorf("postgres user scan: %w", err)
		}
		users = append(users, &user)
	}
	return users, nil
}

func (s *Users) SaveUser(ctx context.Context, user *domain.User) error {
	q := `INSERT INTO users (id, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5);`

	_, err := s.db.ExecContext(ctx, q,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.Created_at)

	if err != nil {
		return fmt.Errorf("postgres save user: %w", err)
	}
	return nil
}

func (s *Users) RemoveUser(ctx context.Context, id string) error {
	q := `DELETE FROM users WHERE id=$1;`

	res, err := s.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("postgres remove user: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("getting affected rows err: %w", err)
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}
