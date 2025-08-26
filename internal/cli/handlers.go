package cli

import (
	"context"
	"fmt"
	"main/internal/domain"
	"main/internal/repo"

	"github.com/fatih/color"
)

func addHandler(ctx context.Context, taskList *domain.TaskList, repo *repo.Repository) {
	title := scanCommand("Task title: ")
	desc := scanCommand("Task description: ")
	t, err := domain.NewTask(title, desc)
	if err != nil {
		fmt.Println("task create error: %w", err)
		return
	}
	if err := taskList.CreateTask(t); err != nil {
		fmt.Println(err)
		pressEnter()
		clear()
		return
	}
	if err := repo.SaveTask(ctx, t); err != nil {
		fmt.Println(err)
		return
	}
	clear()
}

func deleteHandler(ctx context.Context, taskList *domain.TaskList, repo *repo.Repository, filteredList []*domain.Task) {
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
	if err := repo.RemoveTask(ctx, idToRemove); err != nil {
		fmt.Println(err)
		return
	}
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
