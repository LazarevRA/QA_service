FROM golang:1.25.3-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN GOOS=linux GOARCH=amd64 go build -o /app/main ./cmd/api/main.go

EXPOSE 8080

CMD ["/app/main"]