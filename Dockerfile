FROM golang:1.21-bookworm AS builder

WORKDIR /app

# optimize rebuilds
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o goya cmd/server/main.go

FROM golang:1.21-bookworm AS runtime

WORKDIR /app
COPY --from=builder /app/goya goya

EXPOSE 8080
CMD ["./goya"]
