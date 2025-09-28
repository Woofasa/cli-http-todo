package cli

import (
	"context"
	"fmt"
	"main/internal/domain"
	"main/internal/usecase"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func Run(usecase *usecase.App) error {
	handler := &Handler{usecase: usecase}
	ctx := context.Background()

	sort := "default"
	filter := "default"
	descShown := false
	running := true
	for running {
		printHeading(filter, sort)

		loaded, err := usecase.AllTasks(ctx)
		if err != nil {
			return fmt.Errorf("run error: %w", err)
		}

		filtered := usecase.Filter(filter, loaded)
		sorted := usecase.Sort(sort, filtered)
		usecase.Filter(filter, loaded)

		if len(sorted) == 0 {
			fmt.Printf("%s\n", color.HiGreenString("Add some tasks."))
		}
		for i, v := range sorted {
			printTask(i, v, descShown)
		}
		command := askCommand()
		switch strings.ToLower(command) {
		case "exit":
			running = false
			clear()
		case "delete":
			handler.RemoveHandler(ctx, sorted)
		case "add":
			handler.AddHandler(ctx)
		case "close":
			handler.CloseHandler(ctx, sorted)
		case "open":
			handler.OpenHandler(ctx, sorted)
		case "filter":
			filter = askFilter()
		case "sort":
			sort = askSort()
		case "change description":
			handler.ChangeDescriptionHandler(ctx, sorted)
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
		if !task.Status {
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
		if !task.Status {
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
