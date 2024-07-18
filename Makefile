.PHONY: build run test

install:
	go install github.com/onsi/ginkgo/v2/ginkgo
	go mod download

run:
	go run cmd/test-client/test-client.go

test:
	cd test/integration && go test