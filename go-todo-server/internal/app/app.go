package app

import (
	"context"
	"fmt"
	"main/internal/repo"
	"main/internal/repo/postgres"
)

type App struct {
	Repo repo.Storage
}

func NewApp(ctx context.Context) (*App, error) {
	postgres, err := postgres.New()
	if err != nil {
		return nil, fmt.Errorf("postgres new: %w", err)
	}
	if err := postgres.Init(ctx); err != nil {
		return nil, fmt.Errorf("init error: %w", err)
	}

	return &App{
		Repo: postgres,
	}, nil
}
