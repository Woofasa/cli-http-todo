package create_test

import (
	"context"
	"errors"
	"main/internal/domain"
	"main/internal/usecase/task/create"
	"testing"
)

type fakeTaskCreator struct {
	called  bool
	wantErr error
}

func (fc *fakeTaskCreator) SaveTask(ctx context.Context, dto *domain.Task) error {
	fc.called = true
	return fc.wantErr
}

func TestCreator_Create_Success(t *testing.T) {
	fk := &fakeTaskCreator{}
	uc := create.New(fk)
	dto := create.TaskInput{
		Title:       "test",
		Description: "test",
	}
	_, err := uc.Create(context.Background(), dto)
	if err != nil {
		t.Fatalf("Expected no error: %v", err)
	}
	if fk.called != true {
		t.Error("Expected SaveTask calling")
	}
}

func TestCreator_Create_Error(t *testing.T) {
	wantErr := errors.New("something went wrong")
	fk := &fakeTaskCreator{
		wantErr: wantErr,
	}
	uc := create.New(fk)
	dto := create.TaskInput{
		Title:       "test",
		Description: "test",
	}

	_, err := uc.Create(context.Background(), dto)
	if !errors.Is(err, wantErr) {
		t.Fatalf("Expected error %v", wantErr)
	}

	if !fk.called {
		t.Error("Expected SaveTaskCall")
	}
}
