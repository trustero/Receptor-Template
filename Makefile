GOCMD=go
GOTEST=$(GOCMD) test
RECEPTOR_PACKAGE?=./trr-receptorName/receptorPackage

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test-all test coverage goreportcard

all: help

## Test:
test-all: test coverage goreportcard ## Run all tests, this is what Trustero will evaluate

test: ## Run the tests written for the go files of the project
	@echo '${CYAN}Begin go tests${RESET}'
	$(GOTEST) -v -race $(RECEPTOR_PACKAGE)
	@echo '${CYAN}End go tests${RESET}'

coverage: ## Run the tests of the project and report coverage of package functions
	@echo '${CYAN}Begin go code coverage tests${RESET}'
	$(GOTEST) -cover -covermode=count -coverprofile=profile.cov $(RECEPTOR_PACKAGE)
	$(GOCMD) tool cover -func profile.cov
	@echo '${CYAN}End go code coverage tests${RESET}'

goreportcard: ## Run the goreportcard-cli and report the results
	@echo '${CYAN}Begin goreportcard${RESET}'
	goreportcard-cli  -d $(RECEPTOR_PACKAGE) -v
	@echo '${CYAN}End goreportcard${RESET}'


## Help:
help: ## Show this help.
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)