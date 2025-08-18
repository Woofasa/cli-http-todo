package cli

import (
	"fmt"
	"main/internal/domain"
	"main/internal/repo"
	"main/internal/repo/json_storage"
)

type Service struct{}

func Run() error {
	taskList := domain.NewTaskList()

	jsonstorage := json_storage.JSONStorage{}
	repo := repo.Repository{
		DB: map[string]repo.StorageManager{
			"json": jsonstorage,
		},
	}

	repo.LoadAll("json", taskList)
	repo.SaveAll(taskList)

	fmt.Println("Здарова. Текущий список задач: ")
	return nil
}
