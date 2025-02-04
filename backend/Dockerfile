FROM golang:1.23.4-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы в контейнер
COPY . .

COPY ./ssl /etc/ssl/

# Установка swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Генерация документации и сборка
RUN swag init --parseDependency --parseInternal -g cmd/main.go && \
    go build -o main ./cmd

# Устанавливаем значение по умолчанию для переменной
ARG CONFIG
ENV CONFIG_PATH=/app/config/${CONFIG}

# Устанавливаем команду по умолчанию для запуска приложения
CMD ["./main"]
