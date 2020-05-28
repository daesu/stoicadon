.PHONY: generate setup clean test run build testsum

setup:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.23.6

generate:
	cd api/graphql ; go run github.com/99designs/gqlgen generate --verbose

lint:
	golangci-lint run

build:
	GO111MODULE=on go mod tidy

testsum:
	gotestsum -- -p 1 -count=1 -race -cover ./...

run:
	go run cmd/api/main.go