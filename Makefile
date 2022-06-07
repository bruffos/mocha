PROJECT := mocha
MAIN := cmd/hello/main.go

.ONESHELL:
.DEFAULT_GOAL := help

# allow user specific optional overrides
-include Makefile.overrides

export

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## run application
	@go run $(MAIN)

.PHONY: test
test: ## run tests
	@go test -v ./...

.PHONY: bench
bench: ## run benchmarks
	@go test -v ./... -bench=. -count 2 -benchmem -run=^#

.PHONY: coverage
coverage: ## run tests and generate coverage report
	@mkdir -p coverage
	@go test -v ./... -race -coverprofile=coverage/coverage.out -covermode=atomic
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html

vet: ## check go code
	@go vet ./...

fmt: ## run gofmt in all project files
	@go fmt ./...

check: vet ## check source code
	@staticcheck ./...

build: ## build application
	@go build -o bin/hello $(MAIN)

deps: ## check dependencies
	@go mod verify

download: ## download dependencies
	@go mod download

prep: ## prepare local development  environment
	@echo "local tools"
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@npm i --no-package-lock