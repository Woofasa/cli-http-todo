package repo

import (
	"context"
	"fmt"
	"main/internal/domain"
)

type Storage interface {
	GetTaskByID(ctx context.Context, id string) (*domain.Task, error)
	GetTasks(ctx context.Context) (map[string]*domain.Task, error)
	SaveTask(ctx context.Context, tasks *domain.Task) error
	RemoveTask(ctx context.Context, id string) error
	CloseTask(ctx context.Context, id string) error
	OpenTask(ctx context.Context, id string) error
	ChangeDesc(ctx context.Context, newDesc string, id string) error
}

type Repository struct {
	DBs map[string]Storage
}

func (r *Repository) GetTaskByID(ctx context.Context, id string, primaryDB string) (*domain.Task, error) {
	primary := r.DBs[primaryDB]
	result, err := primary.GetTaskByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get tasks error: %w", err)
	}
	return result, nil
}

func (r *Repository) GetTasks(ctx context.Context, primaryDB string) (map[string]*domain.Task, error) {
	primary := r.DBs[primaryDB]
	result, err := primary.GetTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("get tasks error: %w", err)
	}
	return result, nil
}

func (r *Repository) SaveTask(ctx context.Context, task *domain.Task) error {
	for _, db := range r.DBs {
		if err := db.SaveTask(ctx, task); err != nil {
			return fmt.Errorf("save task error: %w", err)
		}
	}
	return nil
}

func (r *Repository) RemoveTask(ctx context.Context, id string) error {
	for _, db := range r.DBs {
		if err := db.RemoveTask(ctx, id); err != nil {
			return fmt.Errorf("remove task error: %w", err)
		}
	}
	return nil
}

func (r *Repository) CloseTask(ctx context.Context, id string) error {
	for _, db := range r.DBs {
		if err := db.CloseTask(ctx, id); err != nil {
			return fmt.Errorf("close task error: %w", err)
		}
	}
	return nil
}

func (r *Repository) OpenTask(ctx context.Context, id string) error {
	for _, db := range r.DBs {
		if err := db.OpenTask(ctx, id); err != nil {
			return fmt.Errorf("open task error: %w", err)
		}
	}
	return nil
}

func (r *Repository) ChangeDesc(ctx context.Context, newDesc string, id string) error {
	for _, db := range r.DBs {
		if err := db.ChangeDesc(ctx, newDesc, id); err != nil {
			return fmt.Errorf("desc change error: %w", err)
		}
	}
	return nil
}
