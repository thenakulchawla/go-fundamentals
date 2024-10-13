.PHONY: build run test lint fmt vet tidy

build:
	go build -o learn-go cmd/main.go

run: build
	./learn-go

test:
	go test ./...

lint:
	golangci-lint run

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

ci: fmt vet lint test build

.DEFAULT_GOAL := build