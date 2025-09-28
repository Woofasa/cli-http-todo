package usecase

import (
	"context"
	"fmt"
	"main/internal/config"
	"main/internal/repo"
	"main/internal/repo/postgres"
)

type App struct {
	TaskStorage repo.TaskStorage
	UserStorage repo.UserStorage
}

func NewApp(ctx context.Context) (*App, error) {
	cfg := config.MustLoad()
	db, err := postgres.NewDB(cfg)
	tasks := postgres.NewTaskStorage(db)
	users := postgres.NewUserStorage(db)
	if err != nil {
		return nil, fmt.Errorf("postgres new: %w", err)
	}
	if err := tasks.Init(ctx); err != nil {
		return nil, fmt.Errorf("init error: %w", err)
	}
	if err := users.Init(ctx); err != nil {
		return nil, fmt.Errorf("init error: %w", err)
	}

	return &App{
		TaskStorage: tasks,
		UserStorage: users,
	}, nil
}
