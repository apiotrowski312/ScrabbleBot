.DEFAULT_GOAL := help
.PHONY: help

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

DOCKER_TEST_COMMAND=docker run --rm -v $(PWD):/gaddag golang:1.13 sh -c "cd /gaddag; go mod download 2> /dev/null;
COLOR_TEST_OUTPUT=sed ''/ok/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

# -----------------------------------------------------------------------------
# Tests
# -----------------------------------------------------------------------------

unit: ## Run unit tests in Docker
	@echo -e ${GREEN}UNIT TESTS${NC}
	@${DOCKER_TEST_COMMAND} go test ./... -cover " | ${COLOR_TEST_OUTPUT}

golden-update: ## update golden files // TODO: Fix perrmision error 
	@echo -e ${GREEN}GOLDENFILES UPDATE${NC}
	@find . -name "*.golden" -type f -delete
	@${DOCKER_TEST_COMMAND} go test ./... -update" | ${COLOR_TEST_OUTPUT}
bench: ## Run benchmark tests in Docker
	@${DOCKER_TEST_COMMAND} go test ./... -bench=. -run=^a"

help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

