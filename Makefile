VERSION := "$(shell git describe --abbrev=0 --tags 2> /dev/null || echo 'v0.0.0')+$(shell git rev-parse --short HEAD 2> /dev/null || echo 'unknown')"

build:
	go build -ldflags "-X main.buildVersion=$(VERSION)"

run:
	go run main.go
