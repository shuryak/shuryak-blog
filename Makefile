GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/user/user.proto
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/articles/articles.proto

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: update
update:
	@go get -u

.PHONY: build-user
build-user:
	@go build -o user cmd/user/*.go

.PHONY: build-articles
build-articles:
	@go build -o articles cmd/articles/*.go
