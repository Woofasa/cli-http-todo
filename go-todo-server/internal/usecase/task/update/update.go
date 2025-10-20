package update

import (
	"context"
	"main/internal/domain"
)

type DTO struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

type Updater interface {
	UpdateTask(ctx context.Context, task *domain.Task) error
}

type UpdateTask struct {
	Updater
}

func New(storage Updater) *UpdateTask {
	return &UpdateTask{
		storage,
	}
}

func (uc *UpdateTask) Execute(ctx context.Context, t *domain.Task, dto DTO) error {
	if dto.Title != nil {
		t.ChangeTitle(*dto.Title)
	}
	if dto.Description != nil {
		t.ChangeDescription(*dto.Description)
	}
	if dto.Status != nil {
		if err := t.ChangeStatus(*dto.Status); err != nil {
			return err
		}
	}
	return uc.UpdateTask(ctx, t)
}
