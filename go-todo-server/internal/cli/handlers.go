package cli

import (
	"context"
	"fmt"
	"main/internal/app"
	"main/internal/domain"
)

type Handler struct {
	App *app.App
}

func (h *Handler) AddHandler(ctx context.Context) {
	dto := app.TaskInput{
		Title:       scanCommand("Task title: "),
		Description: scanCommand("Task description: "),
	}
	h.App.CreateTask(ctx, dto)
	clear()
}

func (h *Handler) RemoveHandler(ctx context.Context, filteredList []*domain.Task) {
	id, err := askID("Task to delete: ", len(filteredList))
	if err != nil {
		fmt.Println(err)
		errorCheck()
		return
	}
	if err := h.App.DeleteTask(ctx, filteredList[id-1].ID); err != nil {
		fmt.Println(err)
		errorCheck()
		return
	}
	clear()
}

func (h *Handler) ChangeDescriptionHandler(ctx context.Context, filteredList []*domain.Task) {
	id, err := askID("Task to change: ", len(filteredList))
	if err != nil {
		fmt.Println(err)
		errorCheck()
		return
	}
	uuid := filteredList[id-1].ID
	newDesc := scanCommand("New description: ")
	if err := h.App.ChangeDescription(ctx, newDesc, uuid); err != nil {
		fmt.Println(err)
		errorCheck()
		return
	}
	clear()
}

func (h *Handler) CloseHandler(ctx context.Context, filteredList []*domain.Task) {
	id, err := askID("Task to close: ", len(filteredList))
	if err != nil {
		fmt.Println(err)
		errorCheck()
		return
	}

	uuid := filteredList[id-1].ID
	if err := h.App.CloseTask(ctx, uuid); err != nil {
		fmt.Println(err)
		errorCheck()
		return
	}
	clear()
}

func (h *Handler) OpenHandler(ctx context.Context, filteredList []*domain.Task) {
	id, err := askID("Task to open: ", len(filteredList))
	if err != nil {
		fmt.Println(err)
		return
	}

	uuid := filteredList[id-1].ID
	if err := h.App.OpenTask(ctx, uuid); err != nil {
		fmt.Println(err)
		errorCheck()
		return
	}

	clear()
}
