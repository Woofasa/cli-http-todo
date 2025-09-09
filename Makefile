.PHONY: backend  frontend dev

backend: 
	cd ./go-todo-server/ && go run ./cmd/main.go
frontend:
	cd ./next-todo-app/ && pnpm dev
dev:
	make -j2 backend frontend