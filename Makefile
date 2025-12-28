.DEFAULT_GOAL := run

.PHONY: run build fmt vet

fmt:
	@go fmt ./...

vet: fmt
	@go vet ./...

build: vet
	@go build -o ./bin/cmd .

run: build
	@go run . Vienna,AT

test: vet
	@go test ./...


