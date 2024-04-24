.DEFAULT_GOAL := build

.PHONY:fmt vet build run test clean
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

run: build
	./main

test:
	go test

clean:
	go clean
