apps = 'api' 'greeter'

VERSION ?= $(shell git rev-parse --short HEAD)-$(shell date -u '+%Y%m%d%I%M%S')


.PHONY: build
build:
	for app in $(apps) ;\
	do \
		CGO_ENABLED=0 go build -o dist/$$app -mod vendor -a -installsuffix cgo -ldflags "-w -s -X hal9000/internal/app/$$app/version.Version=$VERSION" ./cmd/$$app;\
	done

.PHONY: build-image
build-image:
	docker build -t hanamichi/hal9000-api:k8s -f build/api/Dockerfile .
	docker build -t hanamichi/hal9000-srv:k8s -f build/greeter/Dockerfile .
	docker push hanamichi/hal9000-api:k8s
	docker push hanamichi/hal9000-srv:k8s

.PHONY: clean
clean:
	docker rmi hanamichi/hal9000-api:k8s
	docker rmi hanamichi/hal9000-srv:k8s