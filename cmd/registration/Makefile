# ======================================================
# Makefile for registration service (developer use only)
# ======================================================

GO       := go
ROOT_DIR := $(abspath ../..)
BIN_DIR  := $(ROOT_DIR)/bin
SERVICE  := registration

.PHONY: tidy build run

# Tidy dependencies (always run from root)
tidy:
	@echo ">>> Running go mod tidy"
	cd $(ROOT_DIR) && $(GO) mod tidy

# Build only the registration service
build:
	@echo ">>> Building $(SERVICE)"
	$(GO) build -o $(BIN_DIR)/$(SERVICE) .

# Run the registration service directly
run:
	@echo ">>> Running $(SERVICE)"
	$(GO) run .

# Clean the built binary
clean:
	@echo ">>> Cleaning $(SERVICE) binary"
	rm -f $(BIN_DIR)/$(SERVICE)
