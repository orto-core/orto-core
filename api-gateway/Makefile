.PHONY: build

build:
	@go build -o ./build/api-gateway ./cmd/main.go

watch:
	@CompileDaemon -build="make build" -command="./build/api-gateway" -color=true
