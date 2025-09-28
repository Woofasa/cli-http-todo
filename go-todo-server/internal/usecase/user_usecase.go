package usecase

import (
	"context"
	"fmt"
	"main/internal/domain"
)

func (a *App) AllUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := a.UserStorage.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (a *App) CreateUser(ctx context.Context, dto UserInput) (*domain.User, error) {
	u, err := domain.NewUser(dto.Name, dto.Email, dto.Password)
	if err != nil {
		return nil, fmt.Errorf("new user error: %w", err)
	}

	if err := a.UserStorage.SaveUser(ctx, u); err != nil {
		return nil, fmt.Errorf("save user error: %w", err)
	}
	return u, nil
}

func (a *App) RemoveUser(ctx context.Context, id string) error {
	if err := a.UserStorage.RemoveUser(ctx, id); err != nil {
		return err
	}
	return nil
}
