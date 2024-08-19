.PHONY: build test

build:
	@go build -o ./build/auth-service ./cmd/main.go

watch:
	@CompileDaemon -build="make build" -command="./build/auth-service" -color=true


test:
	@go test ./tests -v
