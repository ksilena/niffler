PROJECT_NAME := niffler
SEED := on

.PHONY: build
build:
	CGO_ENABLED=0 go build -o ./${PROJECT_NAME}-cli cmd/cli/main.go

.PHONY: run
run: build
	./${PROJECT_NAME}-cli

.PHONY: lint.install
lint.install:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH}/bin ${LINTER_VERSION}

.PHONY: lint
lint:
	golangci-lint --version
	golangci-lint linters
	golangci-lint run -v

.PHONY: test
test:
	go test -v -race -shuffle=$(SEED) -covermode=atomic -coverprofile=overalls.coverprofile -p 2 ./... 
	go tool cover -func=overalls.coverprofile
	go tool cover -html=overalls.coverprofile -o coverage.html

