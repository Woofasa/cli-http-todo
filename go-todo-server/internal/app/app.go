package app

import (
	"context"
	"fmt"
	"main/internal/repo"
	"main/internal/repo/postgres"
)

type App struct {
	Repo *repo.Repository
}

func NewApp(ctx context.Context) (*App, error) {
	postgres, err := postgres.New()
	if err != nil {
		return nil, fmt.Errorf("postgres new: %w", err)
	}
	if err := postgres.Init(ctx); err != nil {
		return nil, fmt.Errorf("init error: %w", err)
	}

	repo := &repo.Repository{
		DBs: map[string]repo.Storage{
			"postgres": postgres,
		},
	}

	return &App{
		Repo: repo,
	}, nil
}
