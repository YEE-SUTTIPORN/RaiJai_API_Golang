# ใช้ base image ที่มี Go
FROM golang:1.24-alpine

# ติดตั้ง git, curl, bash และ build tools
RUN apk add --no-cache git curl bash

# ติดตั้ง swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

# คัดลอก source code ทั้งหมด
COPY . .

# รัน swag init ตอน build (ใช้ || true กัน error)
RUN swag init || true

# tidy dependencies
RUN go mod tidy

# เปิดพอร์ต
EXPOSE 8080

# ✅ รันแอปแบบปกติ ไม่ใช้ air
CMD ["go", "run", "main.go"]