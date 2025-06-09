# --- Stage 1: Build ---
FROM golang:alpine AS builder

# Install git (optional — needed if you use "go get" or private repos)
RUN apk add --no-cache git

WORKDIR /opt/todo-app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o todo

# --- Stage 2: Runtime ---
FROM alpine:latest

# Create non-root user for security
RUN adduser -D appuser

WORKDIR /home/appuser

# Copy the compiled Go binary
COPY --from=builder /opt/todo-app/todo .

# Use the non-root user
USER appuser

EXPOSE 5000

CMD ["./todo"]
