GOPATH ?= $(shell go env GOPATH)

.PHONY: proto
proto:
	protoc -I=${GOPATH}/src -I=. --go_out=plugins=grpc:. add.proto
