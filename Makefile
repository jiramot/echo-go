BINARY=run

run:
	 go run cmd/main.go

build:
	go build -o bin/engine cmd/main.go

test:	
	go test -short  ./...
	
.PHONY: run