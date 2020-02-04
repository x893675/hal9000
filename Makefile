apps = 'addsrv' 'client'

.PHONY: build
build:
	for app in $(apps) ;\
	do \
		GOOS=linux CGO_ENABLED=0 go build -mod vendor -o dist/$$app -a -installsuffix cgo -ldflags "-w -s" ./cmd/$$app;\
	done

.PHONY: docker
docker:
	docker-compose up --build -d

.PHONY: clean-image
clean-image:
	docker-compose stop
	docker-compose rm -f
	docker images | grep hal9000 | awk '{cmd="docker rmi "$$1":"$$2;system(cmd)}'