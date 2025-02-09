# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/.env .

ENV DB_HOST=localhost \
    DB_USER=your_local_user \
    DB_PASSWORD=your_local_password \
    DB_NAME=your_local_dbname \
    DB_PORT=5432 \
    JWT_SECRET=your_jwt_secret \
    JWT_EXPIRATION=24h

EXPOSE 8080
CMD ["./main"]
