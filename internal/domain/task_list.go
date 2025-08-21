package domain

import (
	"errors"
	"fmt"
	"slices"
)

var (
	ErrAlreadyExists = errors.New("task already exists")
	ErrNotFound      = errors.New("task not found")
	ErrCreatingTask  = errors.New("error creating task")
)

type TaskList struct {
	Tasks map[string]*Task
}

func NewTaskList() *TaskList {
	return &TaskList{
		Tasks: make(map[string]*Task),
	}
}

func Sort(filteredList []*Task) []*Task {
	slices.SortFunc(filteredList, func(a, b *Task) int {
		if a.CreatedAt.Before(b.CreatedAt) {
			return -1
		} else if b.CreatedAt.Before(a.CreatedAt) {
			return 1
		}
		return 0
	})
	return filteredList
}

func (t TaskList) Filter(pattern string) []*Task {
	filtered := make([]*Task, 0, len(t.Tasks))
	switch pattern {
	case "opened":
		for _, v := range t.Tasks {
			if v.Status == Opened {
				filtered = append(filtered, v)
			}
		}
	case "closed":
		for _, v := range t.Tasks {
			if v.Status == Closed {
				filtered = append(filtered, v)
			}
		}
	default:
		for _, v := range t.Tasks {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (t *TaskList) CreateTask(title string, desc string) error {
	task, err := NewTask(title, desc)
	if err != nil {
		return fmt.Errorf("creating task error: %w", err)
	}

	if _, ok := t.Tasks[task.ID]; ok {
		return ErrAlreadyExists
	}

	t.Tasks[task.ID] = task
	return nil
}

func (t *TaskList) RemoveTask(id string) error {
	if _, ok := t.Tasks[id]; !ok {
		return ErrNotFound
	}
	delete(t.Tasks, id)
	return nil
}
