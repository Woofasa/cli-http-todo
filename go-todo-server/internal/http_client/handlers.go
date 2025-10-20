package httpclient

import (
	"context"
	"encoding/json"
	"errors"
	"main/internal/domain"
	"main/internal/usecase"
	"main/internal/usecase/task/create"
	getall "main/internal/usecase/task/getAll"
	"main/internal/usecase/task/gettask"
	"main/internal/usecase/task/remove"
	"main/internal/usecase/task/update"
	"net/http"
)

type Handler struct {
	App *usecase.App
}

func (h *Handler) TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAll := getall.New(h.App.TaskStorage)
		tasks, err := getAll.Execute(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(tasks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
	case http.MethodPost:
		createTask := create.New(h.App.TaskStorage)
		var dto create.TaskInput
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		t, err := createTask.Execute(context.Background(), dto)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidName):
				http.Error(w, "bad request: invalid data", http.StatusBadRequest)
				return
			default:
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
		}
		if err := json.NewEncoder(w).Encode(t); err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		//todo: fix error
	case http.MethodDelete:
		removeTask := remove.New(h.App.TaskStorage)
		id := r.URL.Query().Get("id")
		if err := removeTask.Execute(context.Background(), id); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
		}
	case http.MethodPatch:
		getTask := gettask.New(h.App.TaskStorage)
		updateTask := update.New(h.App.TaskStorage)
		id := r.URL.Query().Get("id")
		task, err := getTask.GetByID(context.Background(), id)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		var dto update.DTO

		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		if err := updateTask.Execute(context.Background(), task, dto); err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(task); err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := h.App.UserStorage.GetUsers(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

	case http.MethodPost:
		var dto usecase.UserInput
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		u, err := h.App.CreateUser(context.Background(), dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(u); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		if err := h.App.UserStorage.RemoveUser(context.Background(), id); err != nil {
			http.Error(w, "bad request. id not found", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
}
