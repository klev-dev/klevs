.PHONY: all build test clean

default: all

all: build test

build:
	go install -v go.klev.dev/klevs/cmd/... 

test:
	go test -cover -timeout 10s ./...

clean: 
	rm -rf ./.direnv/go/bin/
