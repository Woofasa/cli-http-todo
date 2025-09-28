package postgres

import (
	"context"
	"errors"
	"fmt"
	"main/internal/domain"
	"time"

	_ "github.com/lib/pq"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type Tasks struct {
	db *DB
}

func NewTaskStorage(db *DB) *Tasks {
	return &Tasks{
		db: db,
	}
}

func (s *Tasks) Init(ctx context.Context) error {
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
		return fmt.Errorf("cannot create tasks table: %w", err)
	}

	return nil
}

func (s *Tasks) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
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

func (s *Tasks) GetTasks(ctx context.Context) ([]*domain.Task, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT id, title, description, status, created_at, completed_at FROM tasks ORDER BY created_at`)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	result := make([]*domain.Task, 0)
	for rows.Next() {
		var t domain.Task
		var createdAt time.Time
		var completedAt *time.Time

		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &createdAt, &completedAt)
		t.CreatedAt = createdAt
		t.CompletedAt = completedAt

		result = append(result, &t)
	}
	return result, nil
}

func (s *Tasks) SaveTask(ctx context.Context, task *domain.Task) error {
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

func (s *Tasks) UpdateTask(ctx context.Context, task *domain.Task) error {
	q := `UPDATE tasks SET title=$1, description=$2, status=$3, completed_at=$4 WHERE id=$5`
	_, err := s.db.ExecContext(ctx, q,
		task.Title,
		task.Description,
		task.Status,
		task.CompletedAt,
		task.ID)

	if err != nil {
		return fmt.Errorf("postgres update exec: %w", err)
	}
	return nil
}

func (s *Tasks) RemoveTask(ctx context.Context, id string) error {
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

func (s *Tasks) CloseTask(ctx context.Context, id string) error {
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

func (s *Tasks) Filtered(ctx context.Context, pattern string) ([]*domain.Task, error) {
	q := `SELECT id, title, description, status, created_at, completed_at FROM tasks WHERE status=$1;`

	rows, err := s.db.QueryContext(ctx, q, pattern)
	if err != nil {
		return nil, fmt.Errorf("filter postgres error")
	}
	defer rows.Close()

	tasks := make([]*domain.Task, 0)

	for rows.Next() {
		var t domain.Task
		var createdAt time.Time
		var completedAt *time.Time
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &createdAt, &completedAt); err != nil {
			return nil, fmt.Errorf("filter scan error: %w", err)
		}

		t.CreatedAt = createdAt
		t.CompletedAt = completedAt

		tasks = append(tasks, &t)
	}
	return tasks, nil
}
