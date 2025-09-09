package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"main/internal/domain"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type PostgresStorage struct {
	db *sql.DB
}

func New() (*PostgresStorage, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("cant load env file: %v", &err)
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("cannot open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database doesnt answer %w", err)
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) Init(ctx context.Context) error {
	q := `CREATE TABLE IF NOT EXISTS tasks(
		id UUID PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		status BOOLEAN NOT NULL DEFAULT FALSE,
		created_at  TIMESTAMP NOT NULL,
		completed_at TIMESTAMP
	);`

	_, err := s.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create table: %w", err)
	}

	return nil
}

func (s *PostgresStorage) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
	rows := s.db.QueryRowContext(ctx, `SELECT id, title, description, status, created_at, completed_at FROM tasks WHERE id = $1`, id)

	var t domain.Task
	var createdAt time.Time
	var completedAt *time.Time

	if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &createdAt, &completedAt); err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	t.CreatedAt = createdAt
	t.CompletedAt = completedAt

	return &t, nil
}

func (s *PostgresStorage) GetTasks(ctx context.Context) (map[string]*domain.Task, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT id, title, description, status, created_at, completed_at FROM tasks`)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	result := make(map[string]*domain.Task)
	for rows.Next() {
		var t domain.Task
		var createdAt time.Time
		var completedAt *time.Time

		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &createdAt, &completedAt)
		t.CreatedAt = createdAt
		t.CompletedAt = completedAt

		result[t.ID] = &t
	}
	return result, nil
}

func (s *PostgresStorage) SaveTask(ctx context.Context, task *domain.Task) error {
	q := `INSERT INTO tasks (id, title, description, status, created_at, completed_at) VALUES ($1, $2, $3, $4, $5, $6);`

	_, err := s.db.ExecContext(ctx, q,
		task.ID,
		task.Title,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.CompletedAt)

	if err != nil {
		return fmt.Errorf("postgres error: %w", err)
	}
	return nil
}

func (s *PostgresStorage) RemoveTask(ctx context.Context, id string) error {
	q := `DELETE FROM tasks WHERE id = $1;`

	res, err := s.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("postgres error: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("cannot get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (s *PostgresStorage) CloseTask(ctx context.Context, id string) error {
	q := `UPDATE tasks SET status = false, completed_at = NOW() WHERE id = $1 AND status = true;`
	res, err := s.db.ExecContext(ctx, q,
		id)
	if err != nil {
		return fmt.Errorf("postgres error: %w", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("task is already closed: %w", err)
	}

	return nil
}

func (s *PostgresStorage) OpenTask(ctx context.Context, id string) error {
	q := `UPDATE tasks SET status = true, completed_at = NULL WHERE id = $1 AND status = false;`
	res, err := s.db.ExecContext(ctx, q,
		id)
	if err != nil {
		return fmt.Errorf("postgres error: %w", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("task is already opened: %w", err)
	}

	return nil
}

func (s *PostgresStorage) ChangeDesc(ctx context.Context, newDesc string, id string) error {
	q := `UPDATE tasks SET description = $1 WHERE id = $2;`
	_, err := s.db.ExecContext(ctx, q,
		newDesc,
		id)
	if err != nil {
		return fmt.Errorf("postgres error: %w", err)
	}
	return nil
}
