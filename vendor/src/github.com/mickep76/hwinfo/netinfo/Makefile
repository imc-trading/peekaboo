all: test

format:
	gofmt -w=true .

test: format
	golint
	go vet
	go build mem.go mem_linux.go
	go build mem.go mem_darwin.go
