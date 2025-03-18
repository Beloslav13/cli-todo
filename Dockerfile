FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN apk add --no-cache libpq

COPY . .

# Нет компиляции на данном этапе, оставляем все для go run
# RUN CGO_ENABLED=0 GOOS=linux go build -o cli-todo ./cmd/cli/main.go

# Финальный контейнер
#FROM alpine:latest

#WORKDIR /root/

#COPY --from=builder /app .

# Устанавливаем необходимые библиотеки (если нужно для PostgreSQL)
#RUN apk add --no-cache libpq

#ENTRYPOINT ["go", "run", "cmd/cli/main.go"]
