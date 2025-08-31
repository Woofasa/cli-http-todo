package app

import (
	"context"
	"fmt"
	"main/internal/domain"
)

func (a *App) CreateTask(ctx context.Context, title, desc string) error{
	t, err := domain.NewTask(title, desc)
	if err != nil {
		return fmt.Errorf("new task error: %w", err)
	}
	if err := a.TaskList.CreateTask(t); err != nil {
		return fmt.Errorf("create task error: %w", err)
	}
	if err := a.Repo.SaveTask(ctx, t); err != nil {
		return fmt.Errorf("save task error: %w", err)
	}
	return nil
}

func (a *App) DeleteTask(ctx context.Context, uuid string) error{
	if err := a.TaskList.RemoveTask(uuid); err != nil {
		return fmt.Errorf("remove task from list: %w", err)
	}
	if err := a.Repo.RemoveTask(ctx, uuid); err != nil {
		return fmt.Errorf("remove task from storage: %w", err)
	}
	return nil
}

func (a *App) ChangeDescription(ctx context.Context, desc string, uuid string) error{
	if err := a.TaskList.ChangeDescription(uuid, desc); err != nil{
		return fmt.Errorf("change description: %w", err)
	}
	if err := a.Repo.ChangeDesc(ctx, desc, uuid); err != nil {
		return fmt.Errorf("change desc: %w", err)
	}
	return nil
}

func (a *App) CloseTask(ctx context.Context, uuid string) error{
	if err := a.TaskList.CloseTask(uuid); err != nil {
		return fmt.Errorf("close task: %w", err)
	}
	if err := a.Repo.CloseTask(ctx, uuid); err != nil {
		return fmt.Errorf("storage close task: %w", err)
	}
	return nil
}

func (a *App) OpenTask(ctx context.Context, uuid string) error{
	if err := a.TaskList.OpenTask(uuid); err != nil {
		return fmt.Errorf("open task: %w", err)
	}
	if err := a.Repo.OpenTask(ctx, uuid); err != nil {
		return fmt.Errorf("storage open task: %w", err)
	}
	return nil
}

func (a *App) All() []*domain.Task{
	return a.TaskList.All()
}

func (a *App) Sort(pattern string, list []*domain.Task) []*domain.Task{
	return a.TaskList.Sort(pattern, list)
}

func (a *App) Filter(pattern string, list []*domain.Task) []*domain.Task{
	return a.TaskList.Filter(pattern, list)
}