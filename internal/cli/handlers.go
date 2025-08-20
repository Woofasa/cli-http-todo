package cli

import (
	"bufio"
	"fmt"
	"main/internal/domain"
	"main/internal/repo"
	"os"
)

func addHandler(taskList *domain.TaskList, repo *repo.Repository) (title, desc string) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Task title: ")
	reader.Scan()
	title = reader.Text()
	fmt.Print("Task description: ")
	reader.Scan()
	desc = reader.Text()
	if err := taskList.CreateTask(title, desc); err != nil {
		fmt.Println(err)
		pressEnter()
	}
	repo.SaveAll(taskList)
	clear()
	return
}

func deleteHandler(taskList *domain.TaskList, repo *repo.Repository, filteredList []*domain.Task) {
	id, err := askID("Task to delete: ", len(filteredList))
	if err != nil {
		fmt.Println(err)
		return
	}

	idToRemove := filteredList[id-1].ID
	if err := taskList.RemoveTask(idToRemove); err != nil {
		fmt.Println(err)
		pressEnter()
	}
	repo.SaveAll(taskList)
	clear()
}

func closeHandler(taskList *domain.TaskList, repo *repo.Repository, filteredList []*domain.Task) {
	id, err := askID("Task to close: ", len(filteredList))
	if err != nil {
		fmt.Println(err)
		return
	}

	idToRemove := filteredList[id-1].ID
	if err := taskList.Tasks[idToRemove].CloseTask(); err != nil {
		fmt.Println(err)
		pressEnter()
		return
	}
	repo.SaveAll(taskList)
	clear()
}

func openHandler(taskList *domain.TaskList, repo *repo.Repository, filteredList []*domain.Task) {
	id, err := askID("Task to open: ", len(filteredList))
	if err != nil {
		fmt.Println(err)
		return
	}

	idToRemove := filteredList[id-1].ID
	if err := taskList.Tasks[idToRemove].OpenTask(); err != nil {
		fmt.Println(err)
		pressEnter()
		return
	}
	repo.SaveAll(taskList)
	clear()
}
