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
	taskList := domain.NewTaskList()
	jsonstorage := json_storage.JSONStorage{}
	repo := &repo.Repository{
		DB: map[string]repo.StorageManager{
			"json": jsonstorage,
		},
	}

	if err := repo.LoadAll("json", taskList); err != nil {
		log.Fatalf("loaging error: %v", err)
	}

	running := true
	for running {
		defaultSort := taskList.DefaultSort()
		fmt.Println(color.HiRedString("Task list:"))
		for i, v := range defaultSort {
			if v.Status == domain.Closed {
				fmt.Printf("%s: %s %s %s\nCreated at: %s\nClosed at: %s\n\n",
					color.HiRedString((strconv.Itoa(i + 1))),
					color.HiGreenString(v.Title),
					color.HiRedString("-"),
					getStatusString(v),
					v.CreatedAt.Format("02.01.2006 | 15:04"),
					v.CompletedAt.Format("02.01.2006 | 15:04"),
				)
			} else {
				fmt.Printf("%s: %s %s %s\nCreated: %s\n\n",
					color.HiRedString(strconv.Itoa(i+1)),
					color.HiGreenString(v.Title),
					color.HiRedString("-"),
					getStatusString(v),
					v.CreatedAt.Format("02.01.2006 | 15:04"))
			}
		}
		command := askCommand()
		switch strings.ToLower(command) {
		case "exit":
			running = false
		case "delete":
			deleteHandler(taskList, repo, defaultSort)
		case "add":
			addHandler(taskList, repo)
		case "close":
			closeHandler(taskList, repo, defaultSort)
		case "open":
			openHandler(taskList, repo, defaultSort)
		default:
			fmt.Printf("%s\n", color.HiRedString("Unknown command"))
			pressEnter()
		}

	}
	return nil
}
