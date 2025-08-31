package domain

import (
	"errors"
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

func (t TaskList) All() []*Task{
	result := make([]*Task, 0, len(t.Tasks))
	for _, v := range t.Tasks{
		result = append(result, v)
	}
	return result
}

func (t TaskList) Sort(pattern string, taskList []*Task) []*Task {
	switch pattern {
	case "created_at":
		slices.SortFunc(taskList, func(a, b *Task) int {
			if a.CreatedAt.Before(b.CreatedAt) {
				return -1
			} else if b.CreatedAt.Before(a.CreatedAt) {
				return 1
			}
			return 0
		})
	case "name":
		slices.SortFunc(taskList, func(a, b *Task) int {
			if a.Title > b.Title {
				return -1
			} else if b.Title > a.Title {
				return 1
			}
			return 0
		})
	case "completed_at":
		slices.SortFunc(taskList, func(a, b *Task) int {
			if a.CompletedAt.Before(b.CompletedAt) {
				return -1
			} else if b.CompletedAt.Before(a.CompletedAt) {
				return 1
			}
			return 0
		})
	default:
		slices.SortFunc(taskList, func(a, b *Task) int {
			if a.CreatedAt.Before(b.CreatedAt) {
				return -1
			} else if b.CreatedAt.Before(a.CreatedAt) {
				return 1
			}
			return 0
		})
	}
	return taskList
}

func (t TaskList) Filter(pattern string, taskList []*Task) []*Task {
	filtered := make([]*Task, 0, len(taskList))
	switch pattern {
	case "opened":
		for _, v := range taskList {
			if v.Status == Opened {
				filtered = append(filtered, v)
			}
		}
	case "closed":
		for _, v := range taskList {
			if v.Status == Closed {
				filtered = append(filtered, v)
			}
		}
	default:
		return taskList
	}
	return filtered
}

func (t *TaskList) CreateTask(task *Task) error {
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

func (t *TaskList) ChangeDescription(uuid, desc string) error{
	task, ok := t.Tasks[uuid]
	if !ok{
		return ErrNotFound
	}
	task.ChangeDescription(desc)
	return nil
}

func (t *TaskList) CloseTask(uuid string) error{
	task, ok := t.Tasks[uuid]
	if !ok{
		return ErrNotFound
	}
	task.CloseTask()
	return nil
}

func (t *TaskList) OpenTask(uuid string) error{
	task, ok := t.Tasks[uuid]
	if !ok{
		return ErrNotFound
	}
	task.OpenTask()
	return nil
}
