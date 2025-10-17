package gettask

import (
	"context"
	"fmt"
	"main/internal/domain"
)

type IDTaker interface {
	GetByID(ctx context.Context, id string) (*domain.Task, error)
}

type TaskByID struct {
	IDTaker
}

func New(storage IDTaker) *TaskByID {
	return &TaskByID{
		storage,
	}
}

func (uc *TaskByID) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
	t, err := uc.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting task: %w", err)
	}
	return t, nil
}
