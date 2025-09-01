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
func (h *Handler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.App.All(context.Background(), "sqlite")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
