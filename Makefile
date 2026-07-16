.PHONY: test lint fmt proto

test:
	go test ./...

lint:
	golangci-lint run ./...

fmt:
	golangci-lint fmt

proto:
	buf generate
