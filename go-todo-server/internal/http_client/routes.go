package httpclient

import "net/http"

func NewRouter(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.TasksHandler)
	return mux
}
