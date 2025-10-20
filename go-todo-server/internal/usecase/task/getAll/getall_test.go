package getall_test

import (
	"context"
	"errors"
	"main/internal/domain"
	getall "main/internal/usecase/task/getAll"
	"testing"
)

type fakeGetAll struct {
	tasks   []*domain.Task
	called  bool
	wantErr error
}

func (fc *fakeGetAll) GetTasks(ctx context.Context) ([]*domain.Task, error) {
	fc.called = true
	return fc.tasks, fc.wantErr
}

func TestGetAll_Success(t *testing.T) {
	var tasks []*domain.Task
	task1, _ := domain.NewTask("test1", "test1")
	task2, _ := domain.NewTask("test2", "test2")
	tasks = append(tasks, task1, task2)
	fk := &fakeGetAll{tasks: tasks}
	getAll := getall.New(fk)
	ctx := context.Background()

	testTasks, err := getAll.Execute(ctx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if task1 != testTasks[0] || task2 != testTasks[1] {
		t.Error("Tasks should be equal")
	}
	if !fk.called {
		t.Error("getAll is not called")
	}
}

func TestGetAll_Error(t *testing.T) {
	fk := &fakeGetAll{
		wantErr: errors.New("some db error"),
	}
	getall := getall.New(fk)
	ctx := context.Background()

	_, err := getall.Execute(ctx)
	if err == nil {
		t.Error("Expected error got nil")
	}
	if !fk.called {
		t.Error("getAll is not called")
	}
}
