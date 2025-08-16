# Stage 1 — build the Go binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN go build -o server main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
