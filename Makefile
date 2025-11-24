-include Makefile.local.mk

build:
	go build

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o slu-linux-amd64

build-and-scp:
ifndef TO
	$(error TO is undefined, use 'make build-and-scp TO=root@server.example.com:/slu')
endif

	@make build-linux-amd64
	scp slu-linux-amd64 ${TO}

release:
	goreleaser
	rm -rf ./dist
