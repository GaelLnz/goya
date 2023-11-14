APP_NAME=goya

.PHONY: test lint build run

test:
	go test ./...

lint:
	golangci-lint run

build:
	docker build -t ${APP_NAME} .

run: build
	docker run -p 8080:8080 -t ${APP_NAME}

generate:
	go generate ./...
