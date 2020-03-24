.DEFAULT_GOAL := help
.PHONY: help

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# -----------------------------------------------------------------------------
# Tests
# -----------------------------------------------------------------------------

unit: ## Run unit test in Docker
	@echo -e ${GREEN}UNIT TESTS${NC}
	@docker run --rm -v $(PWD):/gaddag golang:1.13 sh -c "cd /gaddag; go mod download 2> /dev/null; go test ./... -cover" | sed ''/ok/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''
	@echo -e ----------------------------

help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

