FROM golang:1.25-alpine

WORKDIR /api

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

RUN go mod tidy

COPY . .

RUN go install github.com/air-verse/air@latest

CMD ["air", ".air.toml"]