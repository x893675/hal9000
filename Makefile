apps = 'api' 'greeter'

.PHONY: build
build:
	for app in $(apps) ;\
	do \
		CGO_ENABLED=0 go build -o dist/$$app -a -installsuffix cgo -ldflags '-w -s' ./cmd/$$app;\
	done

.PHONY: docker
docker:
	docker build -t hanamichi/hal9000-api:latest -f build/api/Dockerfile .
	docker build -t hanamichi/hal9000-srv:latest -f build/greeter/Dockerfile .
	docker push hanamichi/hal9000-api:latest
	docker push hanamichi/hal9000-srv:latest

.PHONY: clean
clean:
	docker rmi hanamichi/hal9000-api:latest
	docker rmi hanamichi/hal9000-srv:latest