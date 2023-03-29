include .env

.PHONY: all
all:
	make start & PID=$$; \
	sleep 1; \
	make migrate-up; \
	trap 'make drop-table; kill $$PID' EXIT; \
	wait

.PHONY: setup
setup:
	@echo "Installing dependencies..."
	@go mod download
	@go mod download github.com/gorilla/mux
	@go mod download github.com/lib/pq
	@echo "Dependencies installed"

.PHONY: start
start:
	go run cmd/main.go

.PHONY: migrate-up
migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

.PHONY: drop-table
drop-table:
	migrate -path migrations -database "$(DB_URL)" drop -f
