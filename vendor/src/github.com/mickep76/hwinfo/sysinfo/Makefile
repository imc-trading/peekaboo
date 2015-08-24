all: test

format:
	gofmt -w=true .

test: format
	golint
	go vet
	go build sys.go sys_linux.go
	go build sys.go sys_darwin.go
