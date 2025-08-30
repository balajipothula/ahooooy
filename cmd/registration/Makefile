# ======================================================
# Root Makefile for ahooooy project
# - CI/CD entrypoint
# - Multi-binary builds for all cmd/* services
# ======================================================

# Go configuration
GO       := go
BIN_DIR  := bin
CMDS     := $(notdir $(wildcard ./cmd/*))

.PHONY: all tidy test build clean $(CMDS)

# Default: build all services
all: tidy test build

# Tidy dependencies
tidy:
	@echo ">>> Running go mod tidy"
	$(GO) mod tidy

# Run tests
test:
	@echo ">>> Running go test"
	$(GO) test ./... -v

# Build all services
build: $(CMDS)

# Per-service build (e.g. make registration)
$(CMDS):
	@echo ">>> Building service: $@"
	$(GO) build -o $(BIN_DIR)/$@ ./cmd/$@

# Clean up
clean:
	@echo ">>> Cleaning binaries"
	rm -rf $(BIN_DIR)/*
