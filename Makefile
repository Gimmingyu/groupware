CMD = cmd
BUILD = build
GEN = gen

ENTRY = main.go

GOCMD = $$(which go)

all: help

## Build
build: vendor ## Build application
	@echo "${YELLOW}Building server example files ...${YELLOW}"
	@GO111MODULE=on \
	CGO_ENABLED=0 \
	GOARCH=amd64 \
	$(GOCMD) build -mod vendor -o $(BUILD)/$(BINARY_NAME) $(CMD)/$(SERVER)/$(ENTRY)
	@echo "${CYAN}Build done${CYAN}"

## Vet:
vet: ## Run `go vet cmd/main.go`
	@$(GOCMD) vet $(CMD)/$(ENTRY)

## Vendor:
vendor: ## Run `go mod vendor`
	@go mod vendor

## Clean:
clean: ## Clear all generated files
	@echo "${YELLOW}Clean generated files...${YELLOW}"
	@rm -rf $(BUILD)/*
	@rm -rf $(GEN)/*
	@echo "${CYAN}Clean done${CYAN}"

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
			if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
			else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
			}' $(MAKEFILE_LIST)


.PHONY: clean all server client vendor help clean vet