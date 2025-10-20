package gettask

import (
	"context"
	"fmt"
	"main/internal/domain"
)

type IDGetter interface {
	GetByID(ctx context.Context, id string) (*domain.Task, error)
}

type TaskByID struct {
	IDGetter
}

func New(storage IDGetter) *TaskByID {
	return &TaskByID{
		storage,
	}
}

func (uc *TaskByID) Execute(ctx context.Context, id string) (*domain.Task, error) {
	t, err := uc.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting task: %w", err)
	}
	return t, nil
}
