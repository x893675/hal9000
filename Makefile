apps = 'addsrv' 'client'

.PHONY: build
build:
	for app in $(apps) ;\
	do \
		GOOS=linux CGO_ENABLED=0 go build -mod vendor -o dist/$$app -a -installsuffix cgo -ldflags "-w -s" ./cmd/$$app;\
	done
