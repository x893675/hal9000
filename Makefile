GOPATH ?= $(shell go env GOPATH)
apps = 'user'

.PHONY: build
build:
	for app in $(apps) ;\
	do \
		CGO_ENABLED=0 go build -mod vendor -o dist/$$app -a -installsuffix cgo -ldflags "-w -s" ./cmd/$$app;\
	done

.PHONY: docker
docker:
	docker-compose up --build -d

.PHONY: clean-image
clean-image:
	docker-compose stop
	docker-compose rm -f
	docker images | grep hal9000 | awk '{cmd="docker rmi "$$1":"$$2;system(cmd)}'

.PHONY: proto
proto:
	protoc -I=${GOPATH}/src \
	-I=${GOPATH}/src/github.com/gogo/protobuf/gogoproto \
	-I=${GOPATH}/src/github.com/gogo/protobuf/types \
	-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
	-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I=${GOPATH}/src/github.com/mwitkow/go-proto-validators/ \
	-I=. \
	--gogo_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:. \
	--grpc-gateway_out=allow_patch_feature=false,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:. \
	--govalidators_out=gogoimport=true,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:. \
    --swagger_out=openapi/ \
    pb/user/user.proto
