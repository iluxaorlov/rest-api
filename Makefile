.PHONY: build
build:
	go build -v ./cmd/api

.DEFAULT_GOAL := build

.PHONY: test
test:
	go test -v -race -timeout 30s ./...