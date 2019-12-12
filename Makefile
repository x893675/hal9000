apps = 'api' 'greeter'

.PHONY: build
build:
	for app in $(apps) ;\
	do \
		CGO_ENABLED=0 go build -o dist/$$app -a -installsuffix cgo -ldflags '-w -s' ./cmd/$$app;\
	done

.PHONY: docker
docker:
	docker build -t hanamichi/hal9000-api:k8s -f build/api/Dockerfile .
	docker build -t hanamichi/hal9000-srv:k8s -f build/greeter/Dockerfile .
	docker push hanamichi/hal9000-api:k8s
	docker push hanamichi/hal9000-srv:k8s

.PHONY: clean
clean:
	docker rmi hanamichi/hal9000-api:k8s
	docker rmi hanamichi/hal9000-srv:k8s