package cli

import (
	"bufio"
	"fmt"
	"log"
	"main/internal/domain"
	"main/internal/repo"
	"main/internal/repo/json_storage"
	"os"
	"os/exec"
	"strings"
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

	if err := repo.LoadAll("json", taskList); err != nil {
		log.Fatalf("loaging error: %v", err)
	}

	running := true
	for running {
		numerated := taskList.NumeratedSort()
		fmt.Println("Task list:")
		for i, v := range numerated {
			fmt.Println(i+1, v.Title)
		}
		var choice string
		fmt.Println("\"exit\" | \"delete\" | \"add\"")
		fmt.Print("Make one of the commands: ")
		fmt.Scan(&choice)
		switch strings.ToLower(choice) {
		case "exit":
			running = false
		case "delete":
			fmt.Print("Task number for deleting: ")
			var number int
			fmt.Scan(&number)
			taskList.RemoveTask(numerated[number].ID)
			repo.SaveAll(taskList)
			clear()
		case "add":
			title, desc := addHandler()
			taskList.CreateTask(title, desc)
			repo.SaveAll(taskList)
			clear()
		default:
			fmt.Println("Unknown command")
		}

	}
	return nil
}

func addHandler() (title, desc string) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Task title: ")
	reader.Scan()
	title = reader.Text()
	fmt.Print("Task description: ")
	reader.Scan()
	desc = reader.Text()
	return
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
