.PHONY: help build run dev lint migrate-up migrate-down format doc swagger

help:
	@echo "Available commands:"
	echo "build - Build the application"
	echo "run - Run the application"
	echo "dev - Run the application in development mode"
	echo "lint - Lint the application"
	echo "format - Format the application"
	echo "doc - Generate Swagger/OpenAPI documentation"
	echo "migrate-up - Apply database migrations"
	echo "migrate-down - Rollback database migrations"


build:
	go build -o bin/app ./cmd/api

build-consumer:
	go build -o bin/consumer ./cmd/consumer

run:
	go run ./cmd/api

run-consumer:
	go run ./cmd/consumer


dev:
	@echo "Starting development server..."
	go run ./cmd/api

lint: format
	golangci-lint run ./...

format:
	@gofmt -l -w .
	@$(shell go env GOPATH)/bin/goimports -w .

doc:
	@which swag > /dev/null 2>&1 || [ -f $(shell go env GOPATH)/bin/swag ] || (echo "Installing swag..." && go install github.com/swaggo/swag/cmd/swag@latest)
	@if which swag > /dev/null 2>&1; then \
		swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal ; \
	else \
		$(shell go env GOPATH)/bin/swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal ; \
	fi

swagger: doc

migrate-up:
	$(shell go env GOPATH)/bin/migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5439/postgres?sslmode=disable" up

migrate-down:
	$(shell go env GOPATH)/bin/migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5439/postgres?sslmode=disable" down


docker-build:
	docker compose -f docker/docker-compose.yml build

docker-up:
	docker compose -f docker/docker-compose.yml up

docker-down:
	docker compose -f docker/docker-compose.yml down
