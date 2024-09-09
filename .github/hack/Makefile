SHELL := /usr/bin/env bash -o errexit -o pipefail -o nounset

#
##@ Available Commands
#

.PHONY: help
help: ## Print this help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make [command]\033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s:\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: lint
lint: ## Run linter
	@golangci-lint version
	golangci-lint run ./...

.PHONY: test
test: ## Run unit tests
	go test ./... -v -coverprofile=coverage.out

.PHONY: test-coverage
test-coverage: test ## Test and open coverage report in browser
	go tool cover -html=coverage.out

.PHONY: build
build: ## Build binaries
	@for CMD in $$(ls cmd); do \
	  GOOS=linux GOARCH=amd64 go build -o "dist/$${CMD}" "cmd/$${CMD}/main.go"; \
	done
