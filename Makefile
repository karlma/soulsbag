export GO111MODULE := on

all: test

mod:
	go mod tidy

test:
	go test -v ./...

clean:
	go clean -testcache
