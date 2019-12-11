apps = 'api' 'greeter'

.PHONY: build
build:
	for app in $(apps) ;\
	do \
		CGO_ENABLED=0 go build -o dist/$$app -a -installsuffix cgo -ldflags '-w -s' ./cmd/$$app;\
	done