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
		numerated := taskList.NumeratedSort()
		fmt.Println("Task list:")
		for i, v := range numerated {
			fmt.Printf("%d: %s - %s\n", i+1, v.Title, getStatusString(v))
		}
		fmt.Println()
		var choice string
		fmt.Println("list commands: \"exit\" | \"delete\" | \"add\"")
		fmt.Println("task commands: \"close\" | \"open\" | \"rename\" | \"change description\"")
		fmt.Print("Make one of the commands: ")
		fmt.Scan(&choice)
		switch strings.ToLower(choice) {
		case "exit":
			running = false
		case "delete":
			deleteHandler(taskList, repo, numerated)
		case "add":
			addHandler(taskList, repo)
		case "close":
			closeHandler(taskList, repo, numerated)
		case "open":
			openHandler(taskList, repo, numerated)
		default:
			fmt.Println("Unknown command")
			pressEnter()
		}

	}
	return nil
}

func getStatusString(task *domain.Task) string {
	if task.Status {
		return "[ ]"
	}
	return "[X]"
}

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
	taskList.RemoveTask(idToRemove)
	repo.SaveAll(taskList)
	pressEnter()
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

func askID(promt string, max int) (int, error) {
	fmt.Print(promt)
	var id int
	fmt.Scan(&id)
	if id < 0 || id > max {
		return 0, fmt.Errorf("invalid id")
	}
	return id, nil
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func pressEnter() {
	fmt.Print("Press Enter to continue.")
	fmt.Scanln()
	clear()
}
