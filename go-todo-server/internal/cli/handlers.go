package cli

import (
	"context"
	"fmt"
	"main/internal/domain"
	"main/internal/usecase"
)

type Handler struct {
	usecase *usecase.App
}

func (h *Handler) AddHandler(ctx context.Context) {
	dto := usecase.TaskInput{
		Title:       scanCommand("Task title: "),
		Description: scanCommand("Task description: "),
	}
	h.usecase.CreateTask(ctx, dto)
	clear()
}

func (h *Handler) RemoveHandler(ctx context.Context, filteredList []*domain.Task) {
	id, err := askID("Task to delete: ", len(filteredList))
	if err != nil {
		fmt.Println(err)
		errorCheck()
		return
	}
	if err := h.usecase.DeleteTask(ctx, filteredList[id-1].ID); err != nil {
		fmt.Println(err)
		errorCheck()
		return
	}
	clear()
}

func (h *Handler) ChangeDescriptionHandler(ctx context.Context, filteredList []*domain.Task) {
	// id, err := askID("Task to change: ", len(filteredList))
	// if err != nil {
	// 	fmt.Println(err)
	// 	errorCheck()
	// 	return
	// }
	// uuid := filteredList[id-1].ID
	// newDesc := scanCommand("New description: ")
	// if err := h.usecase.ChangeDescription(ctx, newDesc, uuid); err != nil {
	// 	fmt.Println(err)
	// 	errorCheck()
	// 	return
	// }
	clear()
}

func (h *Handler) CloseHandler(ctx context.Context, filteredList []*domain.Task) {
	// id, err := askID("Task to close: ", len(filteredList))
	// if err != nil {
	// 	fmt.Println(err)
	// 	errorCheck()
	// 	return
	// }

	// uuid := filteredList[id-1].ID
	// if err := h.usecase.CloseTask(ctx, uuid); err != nil {
	// 	fmt.Println(err)
	// 	errorCheck()
	// 	return
	// }
	clear()
}

func (h *Handler) OpenHandler(ctx context.Context, filteredList []*domain.Task) {
	// id, err := askID("Task to open: ", len(filteredList))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// uuid := filteredList[id-1].ID
	// if err := h.usecase.OpenTask(ctx, uuid); err != nil {
	// 	fmt.Println(err)
	// 	errorCheck()
	// 	return
	// }

	clear()
}
