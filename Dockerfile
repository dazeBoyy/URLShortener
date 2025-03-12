# 1️⃣ Этап сборки (build)
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник (статическая линковка)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o url-shortener ./cmd/main.go

# 2️⃣ Финальный минимальный образ
FROM alpine:latest

WORKDIR /app

# Копируем бинарный файл из builder-этапа
COPY --from=builder /app/url-shortener /app/url-shortener

COPY .env .env
# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["/app/url-shortener"]
