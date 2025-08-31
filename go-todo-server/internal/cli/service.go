package cli

import (
	"context"
	"fmt"
	"main/internal/app"
	"main/internal/domain"
	"strconv"
	"strings"

	"github.com/fatih/color"
)


func Run(app *app.App) error {
	handler := &Handler{App: app}
	ctx := context.Background()

	loaded, err := app.Repo.GetTasks(ctx, "sqlite")
	if err != nil {
		return fmt.Errorf("run error: %w", err)
	}
	app.TaskList.Tasks = loaded

	sort := "default"
	filter := "default"
	descShown := false
	running := true
	for running {
		printHeading(filter, sort)
		defaultList := app.TaskList.All()
		filteredList := app.TaskList.Filter(filter, defaultList)
		sortedTasks := app.TaskList.Sort(sort,filteredList)
		
		if len(sortedTasks) == 0 {
			fmt.Printf("%s\n", color.HiGreenString("Add some tasks."))
		}
		for i, v := range sortedTasks {
			printTask(i, v, descShown)
		}
		command := askCommand()
		switch strings.ToLower(command) {
		case "exit":
			running = false
			clear()
		case "delete":
			handler.RemoveHandler(ctx, sortedTasks)
		case "add":
			handler.AddHandler(ctx)
		case "close":
			handler.CloseHandler(ctx, sortedTasks)
		case "open":
			handler.OpenHandler(ctx, sortedTasks)
		case "filter":
			filter = askFilter()
		case "sort":
			sort = askSort()
		case "change description":
			handler.ChangeDescriptionHandler(ctx, sortedTasks)
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
