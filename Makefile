.DEFAULT_GOAL := help
.PHONY: help

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

DOCKER_TEST_COMMAND=docker run --rm -v $(PWD):/gaddag golang:1.14 sh -c "cd /gaddag; go mod download;
COLOR_TEST_OUTPUT=sed ''/^ok/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''
TEE_COMMAND=tee -a /tmp/results.tmp

_MAKE=$(MAKE) --no-print-directory 
# -----------------------------------------------------------------------------
# Tests
# -----------------------------------------------------------------------------

unit: ## Run unit tests in Docker.
	@echo -e ${GREEN}UNIT TESTS${NC}
	@${DOCKER_TEST_COMMAND} go test ./... -cover" | ${COLOR_TEST_OUTPUT}

golden-update: ## Update golden files.
	@echo -e ${GREEN}GOLDENFILES UPDATE${NC}
	@find . -name "*.golden" -type f -delete
	@${DOCKER_TEST_COMMAND} go test ./... -update" | ${COLOR_TEST_OUTPUT}

bench: ## Run benchmark tests in Docker.
	@echo -e ${GREEN}BENCHMARK TESTS${NC}
	@${DOCKER_TEST_COMMAND} go test ./... -bench=. -run=^a"

game-bench: ## Run example game benchmark tests in Docker.
	@echo -e ${GREEN}BENCHMARK TESTS - GAME ONLY${NC}
	@${DOCKER_TEST_COMMAND} go test ./... -benchtime=50x -bench=Benchmark_Game -run=^a"

# -----------------------------------------------------------------------------
# Run example game
# -----------------------------------------------------------------------------

_game-run-X:
	@${DOCKER_TEST_COMMAND} cd exampleGame; go run main.go -times=$(NUM) -winshot -loglevel=info" | ${TEE_COMMAND}

game-run: ## Run example game
	@rm -f $(PWD)/exampleGame/img/*.png
	@${DOCKER_TEST_COMMAND} cd exampleGame; go run main.go -screenshot -loglevel=panic"

game-10: ## Run 10 example games
	@${_MAKE} _game-run-X NUM=10	

game-100: ## Run 100 example games
	@${_MAKE} _game-run-X NUM=100	

game-run-X: ## Run X example games. Pass NUM=XXX
	@${_MAKE} _game-run-X
	@${_MAKE} game-get-average 
	@${_MAKE} game-get-average-player

game-clean: ## Clean after example game
	@rm -f $(PWD)/exampleGame/img/*.png
	@rm /tmp/results.tmp

# -----------------------------------------------------------------------------
# Get statistics
# -----------------------------------------------------------------------------

game-get-average: ## Get average points from log file
	@cat /tmp/results.tmp | awk '{ sum += $$7 } END { if (NR > 0) print "Average winning points: " sum / NR }'

game-get-average-player: ## Get average player winner
	@cat /tmp/results.tmp | awk '{ sum += $$5 } END { if (NR > 0) print "Average winning player: " sum / NR }'

# -----------------------------------------------------------------------------
# Other
# -----------------------------------------------------------------------------

help: ## Show this help.
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
