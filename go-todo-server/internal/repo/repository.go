package repo

import (
	"context"
	"main/internal/domain"
)

type Storage interface {
	GetTaskByID(ctx context.Context, id string) (*domain.Task, error)
	GetTasks(ctx context.Context) ([]*domain.Task, error)
	SaveTask(ctx context.Context, task *domain.Task) error
	UpdateTask(ctx context.Context, task *domain.Task) error
	RemoveTask(ctx context.Context, id string) error
}
