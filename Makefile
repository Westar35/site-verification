.PHONY: test lint fmt proto mongo-up mongo-down

test:
	go test ./...

lint:
	golangci-lint run ./...

fmt:
	golangci-lint fmt

proto:
	buf generate

mongo-up:
	docker run -d --name site-verification-mongo -p 27017:27017 mongo:7

mongo-down:
	docker rm -f site-verification-mongo
