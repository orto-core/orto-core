FROM golang:1.21.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./config ./config
COPY ./internal ./internal

COPY .env .

RUN go build -o /app/api-gateway ./cmd/main.go

EXPOSE 8081

RUN useradd -m orto
USER orto

ENTRYPOINT ["/app/api-gateway"]
