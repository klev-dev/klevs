.PHONY: all build test clean

default: all

all: build test

build:
	go install -v go.klev.dev/klevs/cmd/... 

test:
	go test -cover -timeout 10s ./...

clean: 
	rm -rf ./.direnv/go/bin/

.PHONY: gen
gen:
	rm -rf ./api/*.pb.go
	protoc --go_out=. --go_opt=module=go.klev.dev/klevs --go-grpc_out=. --go-grpc_opt=module=go.klev.dev/klevs proto/*.proto
	go fmt ./...
