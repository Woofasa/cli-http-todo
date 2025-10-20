package update_test

import (
	"context"
	"errors"
	"main/internal/domain"
	"main/internal/usecase/task/update"
	"testing"
)

type fakeTaskUpdate struct {
	task    *domain.Task
	called  bool
	wantErr error
}

func (fc *fakeTaskUpdate) UpdateTask(ctx context.Context, task *domain.Task) error {
	fc.called = true
	return fc.wantErr
}

func TestTaskUpdate_Success(t *testing.T) {
	ctx := context.Background()
	task, _ := domain.NewTask("test", "test")
	updatedTitle := "updated"
	updatedDesc := "updated"
	dto := update.DTO{
		Title:       &updatedTitle,
		Description: &updatedDesc,
	}
	fk := &fakeTaskUpdate{
		task: task,
	}
	updateTask := update.New(fk)
	if err := updateTask.Execute(ctx, task, dto); err != nil {
		t.Errorf("undexpected error: %v", err)
	}
	if task.Title != updatedTitle || task.Description != updatedDesc {
		t.Error("updated task and DTO are not equal")
	}
	if !fk.called {
		t.Error("taskUpdate is not called")
	}
}

func TestTaskUpdate_Error(t *testing.T) {
	ctx := context.Background()
	task, _ := domain.NewTask("test", "test")
	dto := update.DTO{}
	fk := &fakeTaskUpdate{
		task:    task,
		wantErr: errors.New("updating error"),
	}
	updateTask := update.New(fk)
	if err := updateTask.Execute(ctx, task, dto); err == nil {
		t.Errorf("expected %v got nil", fk.wantErr)
	}
	if !fk.called {
		t.Error("TaskUpdate is not called")
	}
}
