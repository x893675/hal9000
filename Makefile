GOPATH ?= $(shell go env GOPATH)

apps = 'api-server' 'rpctest'

VERSION ?= $(shell git rev-parse --short HEAD)-$(shell date -u '+%Y%m%d%I%M%S')

.PHONY: build
build:
	for app in $(apps) ;\
	do \
		CGO_ENABLED=0 go build -mod vendor -o dist/$$app -a -ldflags "-w -s -X hal9000/pkg/server/version.Version=${VERSION}" ./cmd/$$app;\
	done

.PHONY: swagger
swagger:
	go run tools/doc-gen/main.go --output=swagger-ui/swagger.json

.PHONY: swagger-server
swagger-server:
	go run swagger-ui/swagger.go
