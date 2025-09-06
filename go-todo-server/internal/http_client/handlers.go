package httpclient

import (
	"context"
	"encoding/json"
	"main/internal/app"
	"net/http"
)

type Handler struct {
	App *app.App
}

// GET /
func (h *Handler) TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasks, err := h.App.All(context.Background(), "postgres")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := json.NewEncoder(w).Encode(tasks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var dto app.TaskInput
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if err := h.App.CreateTask(context.Background(), dto); err != nil {
			http.Error(w, "bad request", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
