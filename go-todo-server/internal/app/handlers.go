package app

import (
	"context"
	"fmt"
	"main/internal/domain"
	"slices"
)

func (a *App) CreateTask(ctx context.Context, dto TaskInput) error {
	t, err := domain.NewTask(dto.Title, dto.Description)
	if err != nil {
		return fmt.Errorf("new task error: %w", err)
	}

	if err := a.Repo.SaveTask(ctx, t); err != nil {
		return fmt.Errorf("save task error: %w", err)
	}
	return nil
}

func (a *App) DeleteTask(ctx context.Context, uuid string) error {
	if err := a.Repo.RemoveTask(ctx, uuid); err != nil {
		return fmt.Errorf("remove task from storage: %w", err)
	}
	return nil
}

func (a *App) ChangeDescription(ctx context.Context, desc string, uuid string) error {

	if err := a.Repo.ChangeDesc(ctx, desc, uuid); err != nil {
		return fmt.Errorf("change desc: %w", err)
	}
	return nil
}

func (a *App) CloseTask(ctx context.Context, uuid string) error {
	if err := a.Repo.CloseTask(ctx, uuid); err != nil {
		return fmt.Errorf("storage close task: %w", err)
	}
	return nil
}

func (a *App) OpenTask(ctx context.Context, uuid string) error {
	if err := a.Repo.OpenTask(ctx, uuid); err != nil {
		return fmt.Errorf("storage open task: %w", err)
	}
	return nil
}

func (a *App) All(ctx context.Context, primaryDB string) ([]*domain.Task, error) {
	taskMap, err := a.Repo.GetTasks(ctx, primaryDB)
	taskList := make([]*domain.Task, 0, len(taskMap))
	if err != nil {
		return nil, fmt.Errorf("getting task: %w", err)
	}
	for _, v := range taskMap {
		taskList = append(taskList, v)
	}
	return taskList, nil
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
