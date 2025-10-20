package remove_test

import (
	"context"
	"errors"
	"main/internal/usecase/task/remove"
	"testing"
)

type fakeTaskDeleter struct {
	called  bool
	wantErr error
}

func (fc *fakeTaskDeleter) RemoveTask(ctx context.Context, id string) error {
	fc.called = true
	return fc.wantErr
}

func TestDeleter_Delete_Success(t *testing.T) {
	fk := &fakeTaskDeleter{}
	fakeID := "test"
	ctx := context.Background()
	deleteTask := remove.New(fk)

	if err := deleteTask.Execute(ctx, fakeID); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if fk.called == false {
		t.Error("RemoveTask is not called")
	}
}

func TestDeleter_Delete_Error(t *testing.T) {
	notFoundErr := errors.New("task not found")
	fk := &fakeTaskDeleter{}
	fakeID := "test"
	ctx := context.Background()
	deleteTask := remove.New(fk)
	fk.wantErr = notFoundErr

	if err := deleteTask.Execute(ctx, fakeID); err == nil {
		t.Error("Expected error got nil")
	}

	if fk.called == false {
		t.Error("RemoveTask is not called")
	}
}
