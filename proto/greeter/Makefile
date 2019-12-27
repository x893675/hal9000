GOPATH ?= $(shell go env GOPATH)

.PHONY: proto
proto:
	protoc -I=${GOPATH}/src -I=${GOPATH}/src/github.com/gogo/protobuf/gogoproto -I=. --gogo_out=. --micro_out=. greeter.proto
