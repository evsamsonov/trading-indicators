.PHONY: help lint test doc
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

pre-push: lint test ## Run golang lint and test

lint: ## Run golang lint using docker
	go mod download
	docker run --rm \
		-v $(shell go env GOPATH)/pkg/mod:/go/pkg/mod \
 		-v ${PWD}:/app \
 		-w /app \
	    golangci/golangci-lint:v2.4 \
	    golangci-lint run -v --modules-download-mode=readonly --fix

test: ## Run tests
	go test ./...
