package cli

import (
	"fmt"
	"log"
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

	if err := taskList.CreateTask("Выгуливать кошку", "Надеть намордник обязательно"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Список задач:")
	for i, v := range taskList.NumeratedSort() {
		fmt.Println(i+1, v.Title)
	}

	return nil
}
