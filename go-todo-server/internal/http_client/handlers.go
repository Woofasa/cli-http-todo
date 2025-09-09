package httpclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"main/internal/app"
	"main/internal/domain"
	"net/http"
	"time"
)

type Handler struct {
	App *app.App
}

func (h *Handler) TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasks, err := h.App.All(context.Background(), "postgres")
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
		var dto app.TaskInput
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		t, err := h.App.CreateTask(context.Background(), dto)
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
		id := r.URL.Query().Get("id")
		if err := h.App.DeleteTask(context.Background(), id); err != nil {
			http.Error(w, "something happened", http.StatusInternalServerError)
		}
	case http.MethodPatch:
		id := r.URL.Query().Get("id")
		task, err := h.App.GetTaskByID(context.Background(), id, "postgres")
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		fmt.Println(task.Status)

		var updateErr error
		if task.Status {
			updateErr = h.App.CloseTask(context.Background(), id)
		} else {
			updateErr = h.App.OpenTask(context.Background(), id)
		}
		if updateErr != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		task.Status = !task.Status
		now := time.Now()
		if task.Status {
			task.CompletedAt = nil
		} else {
			task.CompletedAt = &now
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
