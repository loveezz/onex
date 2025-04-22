
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/main .

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/main /app/main

ENV DB_HOST=localhost \
    DB_PORT=5432 \
    DB_USER=postgres \
    DB_NAME=qwe

ARG DB_PASSWORD=123
ENV DB_PASSWORD=$DB_PASSWORD

EXPOSE 8080

# Запускаем приложение
CMD ["/app/main"]