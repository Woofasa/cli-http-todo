package cli

import (
	"bufio"
	"fmt"
	"main/internal/domain"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

func askCommand() string {
	fmt.Println()
	fmt.Println("commands:")
	fmt.Println("\"exit\" | \"delete\" | \"add\" | \"close\" | \"open\" ")
	fmt.Println("\"rename\" | \"change description\" | \"filter\" | \"sort\" | \"show description\"")

	command := scanCommand("Enter the command: ")
	return command
}

func scanCommand(promt string) string {
	fmt.Print(promt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
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
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default: // linux, darwin, etc.
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func pressEnter() {
	fmt.Print("Press Enter to continue.")
	fmt.Scanln()
}
