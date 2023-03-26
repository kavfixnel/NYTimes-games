.PHONY: test
test:
	go test -v ./...

.PHONY: format
format:
	gofmt -d ./