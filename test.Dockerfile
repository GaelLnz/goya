FROM golang:1.21-bookworm

WORKDIR /app

# optimize rebuilds
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go test ./...
