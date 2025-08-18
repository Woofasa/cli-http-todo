package domain

import (
	"errors"
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

func (t *TaskList) CreateTask(title string, desc string) error {
	task, err := NewTask(title, desc)
	if err != nil {
		return ErrCreatingTask
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
