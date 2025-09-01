package app

import (
	"context"
	"fmt"
	"main/internal/repo"
	"main/internal/repo/sqlite"
)

type App struct {
	Repo *repo.Repository
}

func NewApp(ctx context.Context, dbPath string) (*App, error) {
	sqlite, err := sqlite.New(dbPath)
	if err != nil {
		return nil, fmt.Errorf("init error: %w", err)
	}
	sqlite.Init(ctx)

	repo := &repo.Repository{
		DBs: map[string]repo.Storage{
			"sqlite": sqlite,
		},
	}

	return &App{
		Repo: repo,
	}, nil
}
