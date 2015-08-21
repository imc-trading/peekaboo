all: test

format:
	gofmt -w=true .

test: format
	golint
	go vet
	go build
