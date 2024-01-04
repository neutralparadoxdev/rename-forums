# syntax=docker/dockerfile:1

FROM golang:1.21-bookworm

WORKDIR /goforum

COPY go.mod go.sum ./

RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

# RUN ls

RUN CGO_ENABLED=0 GOOS=linux go build -o ./docker-goforum ./cmd/goforum

EXPOSE 3001/tcp

CMD ["./docker-goforum"]
