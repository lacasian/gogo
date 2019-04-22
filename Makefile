VERSION := $(shell git rev-parse --short HEAD)

build:
	go build -ldflags "-X main.buildVersion=$(VERSION)"

run:
	go run main.go
