LOCAL_BIN:=$(CURDIR)/bin
BASE_STACK = docker compose -f docker-compose.yml

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

compose-up: ## Run docker compose stack (app, Redis, Jaeger)
	$(BASE_STACK) up --build && docker compose logs -f
.PHONY: compose-up

compose-down: ## Down docker compose
	$(BASE_STACK) down --remove-orphans
.PHONY: compose-down

deps: ## deps tidy + verify
	go mod tidy && go mod verify
.PHONY: deps

run: ## run the application
	go run ./cmd/main.go
.PHONY: run

format: ## format code with gofumpt
	gofumpt -w .
.PHONY: format

lint: ## check by golangci linter
	golangci-lint run
.PHONY: lint

test: ## run all tests with race detector and coverage
	go clean -testcache && go test -race -coverprofile=coverage.txt ./...
.PHONY: test

mock: ## generate mocks with mockery
	mockery
.PHONY: mock

vuln: ## run govulncheck
	govulncheck ./...
.PHONY: vuln

bin-deps: ## install tools
	GOBIN=$(LOCAL_BIN) go install tool
.PHONY: bin-deps

pre-commit: format lint test ## run pre-commit checks
.PHONY: pre-commit
