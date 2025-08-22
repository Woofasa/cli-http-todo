package cli

import (
	"fmt"
	"log"
	"main/internal/domain"
	"main/internal/repo"
	"main/internal/repo/json_storage"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func Run() error {
	taskMap := domain.NewTaskList()
	jsonstorage := json_storage.JSONStorage{}
	repo := &repo.Repository{
		DB: map[string]repo.StorageManager{
			"json": jsonstorage,
		},
	}

	if err := repo.LoadAll("json", taskMap); err != nil {
		log.Fatalf("loaging error: %v", err)
	}

	sort := "default"
	filter := "default"
	descShown := false
	running := true
	for running {
		printHeading(filter, sort)
		sortedTasks := domain.Sort(sort, taskMap.Filter(filter))
		for i, v := range sortedTasks {
			printTask(i, v, descShown)
		}
		command := askCommand()
		switch strings.ToLower(command) {
		case "exit":
			running = false
		case "delete":
			deleteHandler(taskMap, repo, sortedTasks)
		case "add":
			addHandler(taskMap, repo)
		case "close":
			closeHandler(taskMap, repo, sortedTasks)
		case "open":
			openHandler(taskMap, repo, sortedTasks)
		case "filter":
			filter = filterHandler()
		case "sort":
			sort = sortHandler()
		case "change description":
			changeDescriptionHandler(taskMap, repo, sortedTasks)
		case "show description":
			descShown = showDescription(descShown)
		default:
			fmt.Printf("%s\n", color.HiRedString("Unknown command"))
			pressEnter()
			clear()
		}
	}
	return nil
}

func printHeading(filter, sort string) {
	filterString := color.HiBlueString("filter: ") + filter
	sortString := color.HiBlueString("sort: ") + sort
	fmt.Print(color.HiRedString("Task list:\n%s %s\n\n", filterString, sortString))
}

func printTask(idx int, task *domain.Task, descShown bool) {
	switch descShown {
	case false:
		if task.Status == domain.Closed {
			fmt.Printf("%s: %s %s %s\nCreated at: %s\nClosed at: %s\n\n",
				color.HiRedString((strconv.Itoa(idx + 1))),
				color.HiGreenString(task.Title),
				color.HiRedString("-"),
				getStatusString(task),
				task.CreatedAt.Format("02.01.2006 | 15:04"),
				task.CompletedAt.Format("02.01.2006 | 15:04"),
			)
		} else {
			fmt.Printf("%s: %s %s %s\nCreated: %s\n\n",
				color.HiRedString(strconv.Itoa(idx+1)),
				color.HiGreenString(task.Title),
				color.HiRedString("-"),
				getStatusString(task),
				task.CreatedAt.Format("02.01.2006 | 15:04"))
		}
	case true:
		if task.Status == domain.Closed {
			fmt.Printf("%s: %s %s %s\n%s\nCreated at: %s\nClosed at: %s\n\n",
				color.HiRedString((strconv.Itoa(idx + 1))),
				color.HiGreenString(task.Title),
				color.HiRedString("-"),
				getStatusString(task),
				color.HiCyanString(task.Description),
				task.CreatedAt.Format("02.01.2006 | 15:04"),
				task.CompletedAt.Format("02.01.2006 | 15:04"),
			)
		} else {
			fmt.Printf("%s: %s %s %s\n%s\nCreated: %s\n\n",
				color.HiRedString(strconv.Itoa(idx+1)),
				color.HiGreenString(task.Title),
				color.HiRedString("-"),
				getStatusString(task),
				color.HiCyanString(task.Description),
				task.CreatedAt.Format("02.01.2006 | 15:04"))
		}
	}

}
