default: build

-include Makefile.local.mk

build:
	go build

build-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o slu-linux-amd64 -trimpath -ldflags="-s -w"

build-darwin-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o slu-darwin-arm64 -trimpath -ldflags="-s -w"

build-and-scp:
ifndef TO
	$(error TO is undefined, use 'make build-and-scp TO=root@server.example.com:/slu')
endif

	@make build-linux-amd64
	scp slu-linux-amd64 ${TO}

release:
	goreleaser
	rm -rf ./dist
