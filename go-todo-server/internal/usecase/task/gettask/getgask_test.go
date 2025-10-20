package gettask_test

import (
	"context"
	"errors"
	"main/internal/domain"
	"main/internal/usecase/task/gettask"
	"testing"
)

type fakeTaskGetter struct {
	task    *domain.Task
	called  bool
	wantErr error
}

func (fc *fakeTaskGetter) GetByID(ctx context.Context, id string) (*domain.Task, error) {
	fc.called = true
	return fc.task, fc.wantErr
}

func TestGetTask_GetTask_Success(t *testing.T) {
	task := &domain.Task{
		ID: "Test",
	}
	expectedID := task.ID
	fk := &fakeTaskGetter{
		task: task,
	}
	uc := gettask.New(fk)

	taskByID, err := uc.Execute(context.Background(), expectedID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if taskByID.ID != expectedID {
		t.Errorf("Expected id %s got %s", task.ID, taskByID.ID)
	}
	if !fk.called {
		t.Errorf("Expected GetByID call")
	}
}

func TestGetTask_GetTask_Error(t *testing.T) {
	fk := &fakeTaskGetter{
		wantErr: errors.New("not found"),
	}
	uc := gettask.New(fk)

	_, err := uc.Execute(context.Background(), "")

	if err == nil {
		t.Errorf("Expected error got nil")
	}
	if !fk.called {
		t.Errorf("Expected GetByID call")
	}
}
