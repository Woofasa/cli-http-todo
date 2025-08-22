package cli

import (
	"bufio"
	"fmt"
	"main/internal/domain"
	"main/internal/repo"
	"os"

	"github.com/fatih/color"
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
		clear()
		return
	}
	repo.SaveAll(taskList)
	clear()
	return
}

func deleteHandler(taskList *domain.TaskList, repo *repo.Repository, filteredList []*domain.Task) {
	id, err := askID("Task to delete: ", len(filteredList))
	if err != nil {
		fmt.Println(err)
		pressEnter()
		clear()
		return
	}

	idToRemove := filteredList[id-1].ID
	if err := taskList.RemoveTask(idToRemove); err != nil {
		fmt.Println(err)
		pressEnter()
		clear()
		return
	}
	repo.SaveAll(taskList)
	clear()
}

func changeDescriptionHandler(taskList *domain.TaskList, repo *repo.Repository, filteredList []*domain.Task) {
	id, err := askID("Task to change: ", len(filteredList))
	if err != nil {
		fmt.Println(err)
		pressEnter()
		clear()
		return
	}
	idToChange := filteredList[id-1].ID
	newDesc := scanCommand("New description: ")
	taskList.Tasks[idToChange].ChangeDescription(newDesc)
	repo.SaveAll(taskList)
	clear()
}

func closeHandler(taskList *domain.TaskList, repo *repo.Repository, filteredList []*domain.Task) {
	id, err := askID("Task to close: ", len(filteredList))
	if err != nil {
		fmt.Println(err)
		return
	}

	idToClose := filteredList[id-1].ID
	if err := taskList.Tasks[idToClose].CloseTask(); err != nil {
		fmt.Println(err)
		pressEnter()
		clear()
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
		clear()
		return
	}
	repo.SaveAll(taskList)
	clear()
}

func filterHandler() string {
	newFilter := scanCommand("Enter the new filter type (opened | closed | default): ")
	fmt.Println(newFilter)
	switch newFilter {
	case "opened", "closed", "default":
		clear()
		return newFilter
	default:
		fmt.Printf("%s\n", color.RedString("Unknown filter. New filter is default."))
		pressEnter()
		clear()
		return "default"
	}
}

func sortHandler() string {
	newSort := scanCommand("Enter the new sort type (created_at | completed_at | name): ")
	fmt.Println(newSort)
	switch newSort {
	case "created_at", "completed_at", "name", "default":
		clear()
		return newSort
	default:
		fmt.Printf("%s\n", color.RedString("Unknown sort. New sort is default."))
		pressEnter()
		clear()
		return "default"
	}
}

func showDescription(currentStatus bool) bool {
	clear()
	return !currentStatus
}
