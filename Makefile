release:
	goreleaser
	rm -rf ./dist

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build
