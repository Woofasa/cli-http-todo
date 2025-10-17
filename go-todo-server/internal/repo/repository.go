package repo

import (
	"context"
	"main/internal/domain"
)

type TaskStorage interface {
	GetByID(ctx context.Context, id string) (*domain.Task, error)
	GetTasks(ctx context.Context) ([]*domain.Task, error)
	SaveTask(ctx context.Context, task *domain.Task) error
	UpdateTask(ctx context.Context, task *domain.Task) error
	RemoveTask(ctx context.Context, id string) error
}

type UserStorage interface {
	GetUsers(ctx context.Context) ([]*domain.User, error)
	SaveUser(ctx context.Context, user *domain.User) error
	RemoveUser(ctx context.Context, id string) error
}
