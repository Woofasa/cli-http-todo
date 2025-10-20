package remove

import (
	"context"
	"fmt"
)

type Deleter interface {
	RemoveTask(ctx context.Context, id string) error
}

type DeleteTask struct {
	Deleter
}

func New(storage Deleter) *DeleteTask {
	return &DeleteTask{
		storage,
	}
}

func (uc *DeleteTask) Execute(ctx context.Context, id string) error {
	if err := uc.RemoveTask(ctx, id); err != nil {
		return fmt.Errorf("remove task from storage: %w", err)
	}
	return nil
}
