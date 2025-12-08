BINARY_NAME=bra
MAIN_DIR=./cmd/bra
BUILD_DIR=bin

COLOR_RESET = \033[0m
COLOR_BOLD = \033[1m
COLOR_GREEN=\033[32m
COLOR_CYAN = \033[36m
LINE_BREAK_INDENT = "\n                              " # This is on purpose to align the help text properly

.DEFAULT_GOAL := help

.PHONY: build test clean help

# Help message cheatsheet:
# Use ##@ Section Name to create sections in the help message.
# Use ## Target: Description to document each target.
# Use backticks `like this` to highlight commands or filenames in cyan.
# Use pipe character | to indicate line breaks in long descriptions.

##@ Help

help:  ## Show this help message
	@awk ' \
		BEGIN { \
			FS = ":.*##"; \
			printf "\nUsage:\n  make <$(COLOR_CYAN)target$(COLOR_RESET)> [$(COLOR_CYAN)arg=value$(COLOR_RESET)]\n"; \
		} \
		function colorize_backticks(str) { \
			while (match(str, /`[^`]+`/)) { \
				text = substr(str, RSTART+1, RLENGTH-2); \
				before = substr(str, 1, RSTART-1); \
				after = substr(str, RSTART+RLENGTH); \
				str = before "$(COLOR_CYAN)" text "$(COLOR_RESET)" after; \
			} \
			return str; \
		} \
		function wrap_text(str) { \
			gsub(/\|/, $(LINE_BREAK_INDENT), str); \
			return str; \
		} \
		/^[a-zA-Z_-]+:.*?##/ { \
			printf "  $(COLOR_CYAN)%-24s$(COLOR_RESET) %s\n", $$1, wrap_text(colorize_backticks($$2)); \
		} \
		/^##@/ { \
			printf "\n $(COLOR_BOLD)%s$(COLOR_RESET)\n\n", substr($$0, 5); \
		} \
	' $(MAKEFILE_LIST) ; \
	printf "\n"

##@ Getting Started

build: clean  ## Build the CLI binary
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_DIR)
	@echo "$(COLOR_GREEN)Built $(BUILD_DIR)/$(BINARY_NAME)$(COLOR_RESET)"
	@echo "You can run the CLI using: ./$(BUILD_DIR)/$(BINARY_NAME)"

test:  ## Run all tests
	go test ./tests/... -v

clean:  ## Clean build artifacts
	rm -rf $(BUILD_DIR)
	go clean
