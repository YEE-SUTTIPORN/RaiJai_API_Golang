# Stage 1: Build
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git curl
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN swag init
RUN go build -o main .

# Stage 2: Production
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs
EXPOSE 8080
CMD ["./main"]