# Makefile for Go Fiber + GORM Project

APP_NAME           := main
MAIN_FILE          := ./go/main.go
BUILD_DIR          := /tmp
ARTIFACT           := $(APP_NAME)-$(shell date -u +%Y%m%d%H%M%S)
MODULE_NAME        := github.com/$(shell basename $(shell git remote get-url origin))/go
ARTIFACT_NAME_FILE := .artifact_name_file
WORKDIR            := go

.PHONY: all init tidy fmt vet lint test build compress

# Run all validations + build
all: init tidy fmt vet lint test build

# Initialize Go module if not already present
init:
	@if [ ! -f "$(WORKDIR)/go.mod" ]; then \
		cd $(WORKDIR) && go mod init $(MODULE_NAME); \
	fi

# Dependency cleanup
tidy:
	@cd $(WORKDIR) && go mod tidy

# Format the code
fmt:
	@cd $(WORKDIR) && go fmt ./...

# Static analysis
vet:
	@cd $(WORKDIR) && go vet ./...

# Lint with golangci-lint (install if not available)
lint:
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	@cd $(WORKDIR) && golangci-lint run ./...

# Unit tests
test:
	@cd $(WORKDIR) && go test ./...

# Build the binary with UTC timestamp artifact
build: tidy
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(ARTIFACT) $(MAIN_FILE)
	@echo $(ARTIFACT) > $(ARTIFACT_NAME_FILE)

# Compress binary with UPX (auto-install if missing)
compress: build
	@if ! command -v upx >/dev/null 2>&1; then \
		sudo apt --yes --quiet update && sudo apt --yes install upx; \
	fi
	@upx --best --lzma $(BUILD_DIR)/$(ARTIFACT)
