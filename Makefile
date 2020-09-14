.DEFAULT_GOAL := help
.PHONY: help

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

DOCKER_TEST_COMMAND=docker run --rm -v $(PWD):/gaddag golang:1.14 sh -c "cd /gaddag; go mod download;
COLOR_TEST_OUTPUT=sed ''/^ok/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''
TEE_COMMAND=tee -a /tmp/results.tmp

# -----------------------------------------------------------------------------
# Tests
# -----------------------------------------------------------------------------

unit: ## Run unit tests in Docker
	@echo -e ${GREEN}UNIT TESTS${NC}
	@${DOCKER_TEST_COMMAND} go test ./... -cover" | ${COLOR_TEST_OUTPUT}

golden-update: ## update golden files
	@echo -e ${GREEN}GOLDENFILES UPDATE${NC}
	@find . -name "*.golden" -type f -delete
	@${DOCKER_TEST_COMMAND} go test ./... -update" | ${COLOR_TEST_OUTPUT}

bench: ## Run benchmark tests in Docker
	@echo -e ${GREEN}BENCHMART TESTS${NC}
	@${DOCKER_TEST_COMMAND} go test ./... -bench=. -run=^a"

help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# -----------------------------------------------------------------------------
# Commands for example game
# -----------------------------------------------------------------------------

e-get-average: ## Get average points from log file
	@cat /tmp/results.tmp | awk '{ sum += $$7 } END { if (NR > 0) print "Average winning points: " sum / NR }'

e-get-average-player: ## Get average player winner
	@cat /tmp/results.tmp | awk '{ sum += $$5 } END { if (NR > 0) print "Average winning player: " sum / NR }'

e-clean:
	@rm -f $(PWD)/exampleGame/img/*.png
	@rm /tmp/results.tmp

_e-run-10:
	@${DOCKER_TEST_COMMAND} cd exampleGame; go run main.go -times=10 -loglevel=panic" | ${TEE_COMMAND}

_e-run-100:
	@${DOCKER_TEST_COMMAND} cd exampleGame; go run main.go -times=100 -loglevel=panic" | ${TEE_COMMAND}

_e-run-X:
	@${DOCKER_TEST_COMMAND} cd exampleGame; go run main.go -times=$(NUM) -loglevel=panic" | ${TEE_COMMAND}

e-run: ## Run example game
	@rm -f $(PWD)/exampleGame/img/*.png
	@${DOCKER_TEST_COMMAND} cd exampleGame; go run main.go -screenshot -winshot -loglevel=info"

e-run-10: _e-run-10 e-get-average e-get-average-player ## run 10 example games

e-run-100: _e-run-100 e-get-average e-get-average-player ## run 100 example games

e-run-X: _e-run-X e-get-average e-get-average-player ## run X example games. Pass NUM=XXX
