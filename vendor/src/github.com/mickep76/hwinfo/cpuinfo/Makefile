all: test

format:
	gofmt -w=true .

test: format
	golint
	go vet
	go build cpu.go cpu_linux.go
	go build cpu.go cpu_darwin.go
