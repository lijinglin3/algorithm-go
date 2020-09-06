all: fmt vet lint test

check: fmt vet lint

fmt:
	test `go fmt ./... | wc -l` -eq 0

vet:
	go vet ./...

lint:
	golangci-lint run ./...

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
