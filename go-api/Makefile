.PHONY: help build run dev lint migrate-up migrate-down format

help:
	@echo "Available commands:"
	echo "build - Build the application"
	echo "run - Run the application"
	echo "dev - Run the application in development mode"
	echo "lint - Lint the application"
	echo "format - Format the application"
	echo "migrate-up - Apply database migrations"
	echo "migrate-down - Rollback database migrations"


build:
	go build -o bin/app ./cmd/api


run:
	go run ./cmd/api


dev:
	@echo "Starting development server..."
	go run ./cmd/api

lint: format
	golangci-lint run ./...

format:
	@gofmt -l -w .
	@$(shell go env GOPATH)/bin/goimports -w .


migrate-up:
	$(shell go env GOPATH)/bin/migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5439/postgres?sslmode=disable" up

migrate-down:
	$(shell go env GOPATH)/bin/migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5439/postgres?sslmode=disable" down


docker-up:
	docker compose -f docker/docker-compose.yml up

docker-down:
	docker compose -f docker/docker-compose.yml down
