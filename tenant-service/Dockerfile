FROM golang:1.21.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./config ./config
COPY ./internal ./internal

COPY .env .

RUN go build -o /app/tenant-service ./cmd/main.go

EXPOSE 8083

RUN useradd -m orto
USER orto

ENTRYPOINT ["/app/tenant-service"]
