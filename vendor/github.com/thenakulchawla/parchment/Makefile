.PHONY: test fmt vet lint

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

check: fmt vet lint test