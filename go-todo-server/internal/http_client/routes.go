package httpclient

import "net/http"

func NewRouter(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", h.TasksHandler)
	mux.HandleFunc("/users", h.UserHandler)
	return mux
}
