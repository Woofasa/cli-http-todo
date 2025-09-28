package create_test

import (
	"context"
	"errors"
	"testing"

	"main/internal/domain"
	create "main/internal/usecase/task"
)

type fakeStorage struct {
	called bool
	saved  *domain.Task
	err    error
}

func (f *fakeStorage) SaveTask(ctx context.Context, task *domain.Task) error {
	f.called = true
	f.saved = task
	return f.err
}

func TestCreator_Create(t *testing.T) {
	cases := []struct {
		name       string
		input      create.TaskInput
		storageErr error
		wantErr    bool
		wantCall   bool
	}{
		{
			name: "success",
			input: create.TaskInput{
				Title:       "Test title",
				Description: "Test description",
			},
			storageErr: nil,
			wantErr:    false,
			wantCall:   true,
		},
		{
			name: "invalid input (empty title)",
			input: create.TaskInput{
				Title:       "",
				Description: "desc",
			},
			storageErr: nil,
			wantErr:    true,
			wantCall:   false, // не должен дернуть SaveTask
		},
		{
			name: "storage error",
			input: create.TaskInput{
				Title:       "Valid title",
				Description: "Valid desc",
			},
			storageErr: errors.New("db error"),
			wantErr:    true,
			wantCall:   true,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			storage := &fakeStorage{err: tc.storageErr}
			uc := create.New(storage)

			task, err := uc.Create(context.Background(), tc.input)

			if (err != nil) != tc.wantErr {
				t.Errorf("expected error=%v, got err=%v", tc.wantErr, err)
			}

			if storage.called != tc.wantCall {
				t.Errorf("expected SaveTask called=%v, got %v", tc.wantCall, storage.called)
			}

			if !tc.wantErr {
				if task == nil {
					t.Fatal("expected task, got nil")
				}
				if task.Title != tc.input.Title {
					t.Errorf("expected title %q, got %q", tc.input.Title, task.Title)
				}
				if task.Description != tc.input.Description {
					t.Errorf("expected description %q, got %q", tc.input.Description, task.Description)
				}
				if storage.saved != task {
					t.Error("expected saved task to be the same instance")
				}
			}
		})
	}
}
