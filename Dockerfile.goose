FROM golang:1.25.3-alpine3.22

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY migrations /migrations

WORKDIR /migrations

CMD ["goose", "-version"]