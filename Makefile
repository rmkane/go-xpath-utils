# Commands
GO  ?= go
FMT ?= goimports-reviser

# ANSI escape codes
RESET   := \033[0m
BOLD    := \033[1m
RED     := \033[91m
GREEN   := \033[92m
YELLOW  := \033[93m
BLUE    := \033[94m
MAGENTA := \033[95m
CYAN    := \033[96m

.PHONY: all help setup tidy update format lint test test-verbose test-with-coverage cleanup

all: tidy format lint test # [Default] Formats, lints, and runs tests

help: # Display this help message
	@printf "\n$(BOLD)Usage:$(RESET) make $(CYAN)[target]$(RESET)\n\n"
	@printf "$(BOLD)Available targets:$(RESET)\n"
	@awk 'BEGIN {FS = ":.*?#"} /^[a-zA-Z_-]+:.*?#/ {printf "  $(CYAN)%-18s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

setup: # Install dev tools
	@echo -e "> $(GREEN)Installing dev tools...$(RESET)"
	@# Version v3.9.0 changes the min Go version to 1.24 (from 1.18)
	@$(GO) install -v github.com/incu6us/goimports-reviser/v3@v3.8.2

tidy: # Tidy the mod file
	@echo -e "> $(GREEN)Tidying Go modules...$(RESET)"
	@$(GO) mod tidy

update: # Update dependencies
	@echo -e "> $(GREEN)Updating dependencies...$(RESET)"
	@$(GO) get -u ./...
	@$(GO) mod tidy

format: # Format the code
	@echo -e "> $(GREEN)Formatting code...$(RESET)"
	@if command -v $(FMT) >/dev/null 2>&1; then \
		$(FMT) -rm-unused -format ./... 2>/dev/null; \
	else \
		echo -e "$(YELLOW)Warning:$(RESET) $(FMT) is not installed. Falling back to [$(BLUE)go fmt$(RESET)]."; \
		echo -e "Run [$(CYAN)make setup$(RESET)] to install it."; \
		$(GO) fmt ./... >/dev/null; \
	fi

lint: # Run the linter
	@echo -e "> $(GREEN)Linting code...$(RESET)"
	@$(GO) vet ./...

test: # Run tests, showing output only if they fail
	@echo -e "> $(GREEN)Running tests...$(RESET)"
	@$(GO) test ./... > /dev/null || $(GO) test ./...

test-verbose: # Run tests in verbose mode
	@echo -e "> $(GREEN)Running tests in verbose mode...$(RESET)"
	@$(GO) test -v ./...

test-with-coverage: # Run tests with coverage
	@echo -e "> $(GREEN)Running tests with coverage...$(RESET)"
	@$(GO) test -v -cover ./...

cleanup: # Remove dev tools
	@echo -e "> $(GREEN)Removing dev tools...$(RESET)"
	@BIN_PATH="$(shell $(GO) env GOBIN)"; \
	[ -z "$$BIN_PATH" ] && BIN_PATH="$(shell $(GO) env GOPATH)/bin"; \
	$(RM) -f "$$BIN_PATH/$(FMT)"
