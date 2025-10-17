package domain_test

import (
	"main/internal/domain"
	"testing"
)

func Test_NewTask(t *testing.T) {
	cases := []struct {
		name    string
		title   string
		desc    string
		wantErr bool
	}{
		{
			name:    "Succes",
			title:   "Test",
			desc:    "Test",
			wantErr: false,
		},
		{
			name:    "Empty title",
			title:   "",
			desc:    "Test",
			wantErr: true,
		},
		{
			name:    "Too long title",
			title:   "TestTestTestTestTestTestTest",
			desc:    "Test",
			wantErr: true,
		},
		{
			name:    "Empty desc",
			title:   "Test",
			desc:    "",
			wantErr: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			task, err := domain.NewTask(tt.title, tt.desc)
			if tt.wantErr {
				if err == nil {
					t.Errorf("%s: expected error, got nil: %v", tt.name, err)
				}
				return
			}
			if err != nil {
				t.Errorf("%s: unexpected error: %v", tt.name, err)
			}
			if task.Title != tt.title {
				t.Errorf("%s: expected title %s but got %s", tt.name, tt.title, task.Title)
			}

			if tt.desc == "" && task.Description != "-" {
				t.Errorf("%s: expected \"-\" in description", tt.name)
			}
		})
	}
}
