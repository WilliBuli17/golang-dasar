FROM golang:alpine

WORKDIR /app

COPY . .

ENV APP_FILE=main.go

CMD ["sh", "-c", "go run ${APP_FILE}"]