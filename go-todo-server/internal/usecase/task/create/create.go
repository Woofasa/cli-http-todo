package create

import (
	"context"
	"fmt"
	"main/internal/domain"
)

type TaskInput struct {
	Title       string
	Description string
}

type TaskCreator interface {
	SaveTask(ctx context.Context, task *domain.Task) error
}

type CreateTask struct {
	TaskCreator
}

func New(storage TaskCreator) *CreateTask {
	return &CreateTask{
		storage,
	}
}

func (uc *CreateTask) Execute(ctx context.Context, dto TaskInput) (*domain.Task, error) {
	t, err := domain.NewTask(dto.Title, dto.Description)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if err := uc.SaveTask(ctx, t); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return t, nil
}
