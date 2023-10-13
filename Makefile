export GO111MODULE := on

all: test

mod:
	go mod tidy

test:
	go test -v --skip TestSoulsBagWatch ./...

test_watch:
	go test -v --run TestSoulsBagWatch

clean:
	go clean -testcache
