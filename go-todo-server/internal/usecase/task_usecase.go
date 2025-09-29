package usecase

import (
	"context"
	"fmt"
	"main/internal/domain"
	"slices"
)

func (a *App) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
	t, err := a.TaskStorage.GetTaskByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting task: %w", err)
	}
	return t, nil
}

func (a *App) CreateTask(ctx context.Context, dto TaskInput) (*domain.Task, error) {
	t, err := domain.NewTask(dto.Title, dto.Description)
	if err != nil {
		return nil, fmt.Errorf("new task error: %w", err)
	}

	if err := a.TaskStorage.SaveTask(ctx, t); err != nil {
		return nil, fmt.Errorf("save task error: %w", err)
	}
	return t, nil
}

func (a *App) DeleteTask(ctx context.Context, uuid string) error {
	if err := a.TaskStorage.RemoveTask(ctx, uuid); err != nil {
		return fmt.Errorf("remove task from storage: %w", err)
	}
	return nil
}

func (a *App) AllTasks(ctx context.Context) ([]*domain.Task, error) {
	taskMap, err := a.TaskStorage.GetTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("app getting tasks: %w", err)
	}
	return taskMap, nil
}

func (a *App) Sort(pattern string, list []*domain.Task) []*domain.Task {
	switch pattern {
	case "created_at":
		slices.SortFunc(list, func(a, b *domain.Task) int {
			if a.CreatedAt.Before(b.CreatedAt) {
				return -1
			} else if b.CreatedAt.Before(a.CreatedAt) {
				return 1
			}
			return 0
		})
	case "name":
		slices.SortFunc(list, func(a, b *domain.Task) int {
			if a.Title > b.Title {
				return 1
			} else if b.Title > a.Title {
				return -1
			}
			return 0
		})
	case "completed_at":
		slices.SortFunc(list, func(a, b *domain.Task) int {
			if a.CompletedAt != nil && b.CompletedAt == nil {
				return -1
			}
			if b.CompletedAt != nil && a.CompletedAt == nil {
				return -1
			}
			if a.CompletedAt == nil && b.CompletedAt == nil {
				return 0
			}
			if a.CompletedAt.Before(*b.CompletedAt) {
				return -1
			}
			if b.CompletedAt.Before(*a.CompletedAt) {
				return -1
			}
			return 0
		})
	default:
		slices.SortFunc(list, func(a, b *domain.Task) int {
			if a.CreatedAt.Before(b.CreatedAt) {
				return -1
			} else if b.CreatedAt.Before(a.CreatedAt) {
				return 1
			}
			return 0
		})
	}
	return list
}

func (a *App) Filter(pattern string, list []*domain.Task) []*domain.Task {
	filtered := make([]*domain.Task, 0, len(list))
	switch pattern {
	case "opened":
		for _, v := range list {
			if v.Status {
				filtered = append(filtered, v)
			}
		}
	case "closed":
		for _, v := range list {
			if !v.Status {
				filtered = append(filtered, v)
			}
		}
	default:
		return list
	}
	return filtered
}

func (a *App) UpdateTask(ctx context.Context, id string, dto UpdateTaskDTO) error {
	t, err := a.GetTaskByID(ctx, id)
	if err != nil {
		return err
	}
	if dto.Title != nil {
		t.Title = *dto.Title
	}
	if dto.Description != nil {
		t.Description = *dto.Description
	}
	if dto.Status != nil {
		if err := t.ChangeStatus(*dto.Status); err != nil {
			return err
		}
	}

	return a.TaskStorage.UpdateTask(ctx, t)
}
