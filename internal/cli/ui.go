package cli

import (
	"fmt"
	"main/internal/domain"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func askCommand() string {
	fmt.Println()
	var command string
	fmt.Println("commands:")
	fmt.Println("\"exit\" | \"delete\" | \"add\"")
	fmt.Println("\"close\" | \"open\" | \"rename\" | \"change description\"")
	fmt.Print("Make one of the commands: ")
	fmt.Scan(&command)
	return command
}

func getStatusString(task *domain.Task) string {
	if task.Status {
		return "[ ]"
	}
	return fmt.Sprintf("[%s]", color.GreenString("X"))
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
