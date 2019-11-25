.PHONY: default_target all build install test coverage lint help

BIN_DIR=$(GOPATH)/bin
LINTER_INSTALLED := $(shell sh -c 'which golangci-lint')

default_target: all

all: lint test build

test: ## Runs application tests
	go test -race -v ./... -coverprofile=coverage.out -covermode=atomic

coverage: test ## Displays test coverage report
	go tool cover -html=coverage.out

lint: ## Runs the go code linter
ifdef LINTER_INSTALLED
	golangci-lint run
else
	$(error golangci-lint not found, skipping linting. Installation instructions: https://github.com/golangci/golangci-lint#ci-installation)
endif

build:
	go build -ldflags="-X main.gitSHA=${GIT_SHA}" cmd/fireplace.go

ifeq ($(shell uname), Linux)
set-service: # Sets the service into Systemd
	$(shell ln -s $(CURDIR)/fireplace.service /etc/systemd/system/fireplace.service && systemctl daemon-reload)
install: set-service # Installs the fireplace server
endif
install:
	$(shell GOBIN=/usr/local/bin go install cmd/fireplace.go)

help: ## Prints this help command
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) |\
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
