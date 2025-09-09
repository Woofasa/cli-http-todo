package httpclient

import (
	"fmt"
	"log"
	"main/internal/http_client/middleware"
	"net/http"
)

func RunServer(h *Handler, addr string) {
	mux := NewRouter(h)

	fmt.Printf("Server is tarting at port %s\n", addr)
	if err := http.ListenAndServe(addr, middleware.WithCors(mux)); err != nil {
		log.Fatal("server failed")
	}
}
