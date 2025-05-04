FROM golang:1.24.1-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && \
    go get github.com/go-chi/cors && \
    go mod tidy
COPY . .

RUN go build -o main .

EXPOSE 9000

CMD ["./main"]