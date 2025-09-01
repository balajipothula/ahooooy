# ======================================================
# Makefile for registration service (developer use only)
# ======================================================

GO       := go
ROOT_DIR := $(abspath ../..)
BIN_DIR  := $(ROOT_DIR)/bin
SERVICE  := registration

.PHONY: all tidy fmt vet build compress run

# Run all validations + build
all: tidy fmt vet build compress run

# Tidy dependencies (always run from root)
tidy:
	@echo ">>> Running `go mod tidy`"
	cd $(ROOT_DIR) && $(GO) mod tidy

# Format the code
fmt:
	@$(GO) fmt ./...

# Static analysis
vet:
	@$(GO) vet ./...

# Build only the registration service
build: tidy
	@echo ">>> Building $(SERVICE)"
	$(GO) build -o $(BIN_DIR)/$(SERVICE) .

# Compress binary with UPX (auto-install if missing)
compress: build
	@if ! command -v upx >/dev/null 2>&1; then \
		sudo apt --yes --quiet update && sudo apt --yes install upx; \
	fi
	@upx --best --lzma $(BIN_DIR)/$(SERVICE)

# Run the registration service directly
run:
	@echo ">>> Running $(SERVICE)"
	$(GO) run .

# Clean the built binary
clean:
	@echo ">>> Cleaning $(SERVICE) binary"
	rm -f $(BIN_DIR)/$(SERVICE)
