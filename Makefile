-include Makefile.local.mk

build:
	go build

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o slu-linux-amd64

release:
	goreleaser
	rm -rf ./dist
