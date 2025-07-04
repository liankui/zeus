.PHONY: env test build

GOPROXY := https://goproxy.cn,direct
export GOPROXY

default: test

env:
	@go version
test: env
	go test -race -covermode=atomic -coverprofile=coverage.out.tmp -coverpkg ./... ./...
	@cat coverage.out.tmp | grep -v "mock_" > coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@go tool cover -func=coverage.out -o coverage.txt
	@tail -n 1 coverage.txt
build: env
	@cd cmd && ./xbuild.sh

# grpcui -plaintext -import-path . -import-path ~/go/src/github.com/googleapis -proto auth.proto localhost:21102
