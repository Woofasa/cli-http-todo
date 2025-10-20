package getall

import (
	"context"
	"fmt"
	"main/internal/domain"
)

type TasksGetter interface {
	GetTasks(ctx context.Context) ([]*domain.Task, error)
}

type GetAll struct {
	TasksGetter
}

func New(storage TasksGetter) *GetAll {
	return &GetAll{
		storage,
	}
}

func (uc *GetAll) Execute(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := uc.GetTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("getall execute: %w", err)
	}
	return tasks, nil
}
