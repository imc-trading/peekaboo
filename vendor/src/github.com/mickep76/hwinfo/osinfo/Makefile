all: test

format:
	gofmt -w=true .

test: format
	golint
	go vet
	go build os.go os_linux.go
	go build os.go os_darwin.go
