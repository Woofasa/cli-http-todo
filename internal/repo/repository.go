package repo

import (
	"errors"
	"main/internal/domain"
)

type StorageManager interface {
	Save(tasks *domain.TaskList) error
	Load(tasks *domain.TaskList) error
}

type Repository struct {
	DB map[string]StorageManager
}

func (r *Repository) SaveAll(tasks *domain.TaskList) error {
	for _, db := range r.DB {
		if err := db.Save(tasks); err != nil {
			return errors.New("save all error")
		}
	}
	return nil
}

func (r Repository) LoadAll(primaryDB string, tasks *domain.TaskList) error {
	if err := r.DB[primaryDB].Load(tasks); err != nil {
		return errors.New("load all error")
	}
	return nil
}
