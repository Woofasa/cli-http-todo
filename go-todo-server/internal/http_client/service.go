package httpclient

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer(h *Handler, addr string) {
	mux := NewRouter(h)

	fmt.Printf("Server is tarting at port %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("server failed")
	}
}
