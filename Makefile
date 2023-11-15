APP_NAME=goya
PORT=8080

.PHONY: test lint build run

download:
	go mod download

test: download
	go test ./...

lint:
	golangci-lint run

build:
	docker build -t ${APP_NAME} .

run: build
	docker run -p ${PORT}:${PORT} -t ${APP_NAME}

generate:
	go generate ./...
