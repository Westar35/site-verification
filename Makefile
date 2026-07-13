.PHONY: build run test lint fmt proto docker-build mongo-up mongo-down

BINARY := dns-check-service

build:
	go build -o bin/$(BINARY) ./cmd

run:
	go run ./cmd

test:
	go test ./...

lint:
	golangci-lint run ./...

fmt:
	gofmt -w .

proto:
	buf generate

docker-build:
	docker build -t $(BINARY) .

mongo-up:
	docker run -d --name dns-check-mongo -p 27017:27017 mongo:7

mongo-down:
	docker rm -f dns-check-mongo
