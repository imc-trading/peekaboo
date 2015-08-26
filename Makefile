all: build

clean:
	rm -rf pkg bin

test: clean
	gb test

build: test
	gb build all

update:
	gb vendor update --all

hwinfo:
	gb vendor update github.com/mickep76/hwinfo
