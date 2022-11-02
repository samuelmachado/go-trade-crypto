.PHONY: build
build:
	@go build -o ./bin/worker  ./cmd/worker

.PHONY: run
run:
	@go run cmd/worker/main.go

