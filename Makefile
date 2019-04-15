all: fmt vet test

fmt:
	test `go fmt ./... | wc -l` -eq 0

vet:
	go vet ./...

test:
	go test -v -race ./...
