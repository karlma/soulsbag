export GO111MODULE := on

all: test

mod:
	go mod tidy

test: mod
	go test -v ./...

