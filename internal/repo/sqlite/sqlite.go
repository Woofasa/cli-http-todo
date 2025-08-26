package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"main/internal/domain"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	db *sql.DB
}

func New(path string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("cannot open database: %w", err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database doesnt answer %w", err)
	}

	return &SQLiteStorage{
		db: db,
	}, nil
}

func (s *SQLiteStorage) Init(ctx context.Context) error {
	q := `CREATE TABLE IF NOT EXISTS tasks(
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		status BOOLEAN NOT NULL DEFAULT 0,
		created_at  TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now')),
		completed_at  TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now'))
	);`

	_, err := s.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create table: %w", err)
	}

	return nil
}

func (s *SQLiteStorage) GetTasks(ctx context.Context) (map[string]*domain.Task, error) {
	q := `SELECT * FROM tasks;`
	result := make(map[string]*domain.Task, 0)

	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("task loading error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var t domain.Task
		var createdAtStr, completedAtStr string
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &createdAtStr, &completedAtStr); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		t.CreatedAt, _ = time.Parse(time.RFC3339Nano, createdAtStr)
		t.CompletedAt, _ = time.Parse(time.RFC3339Nano, completedAtStr)
		result[t.ID] = &t
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return result, nil
}

func (s *SQLiteStorage) SaveTask(ctx context.Context, task *domain.Task) error {
	q := `INSERT INTO tasks (id, title, description, status, created_at, completed_at) VALUES (?, ?, ?, ?, ?, ?);`

	_, err := s.db.ExecContext(ctx, q,
		task.ID,
		task.Title,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.CompletedAt)

	if err != nil {
		return fmt.Errorf("task saving error: %w", err)
	}
	return nil
}

func (s *SQLiteStorage) RemoveTask(ctx context.Context, id string) error {
	q := `DELETE FROM tasks WHERE id = ?;`

	_, err := s.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("removing error: %w", err)
	}

	return nil
}
