all: fmt vet golangci-lint golint test

check: fmt vet golangci-lint golint

fmt:
	test `go fmt ./... | wc -l` -eq 0

vet:
	go vet ./...

golangci-lint:
	golangci-lint run ./...

golint:
	test `golint ./... | grep -v 'should not use dot imports' | wc -l` -eq 0

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
