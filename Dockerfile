# Use the official Golang image as a builder
FROM golang:1.25-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o server .

# Use a minimal image for running the binary
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Copy docs.html if you have it
COPY --from=builder /app/docs.html ./docs.html

# Expose port 8080
EXPOSE 8080

# Run the server
CMD ["./server"]
