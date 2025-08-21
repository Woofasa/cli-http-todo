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

	filter := "default"
	running := true
	for running {
		sortedTasks := domain.Sort(taskMap.Filter(filter))
		fmt.Println(color.HiRedString("Task list:"))
		for i, v := range sortedTasks {
			printTask(i, v)
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
		default:
			fmt.Printf("%s\n", color.HiRedString("Unknown command"))
			pressEnter()
			clear()
		}

	}
	return nil
}

func printTask(idx int, task *domain.Task) {
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
}
