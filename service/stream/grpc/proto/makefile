
OPTIONS=--go_out=. \
        --go_opt=paths=source_relative \
        --go-grpc_out=. \
        --go-grpc_opt=paths=source_relative \
        --validate_out=paths=source_relative,lang=go:.

export PATH := $(PATH):$(shell go env GOPATH)/bin

all:
	@protoc $(OPTIONS) stream.proto

toolinstall:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/envoyproxy/protoc-gen-validate@latest
