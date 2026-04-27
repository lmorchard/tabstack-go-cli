.PHONY: setup build run clean lint format test vendor vendor-check

# Build variables
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION_PKG := github.com/lmorchard/tabstack-go-cli/cmd
LDFLAGS := -X $(VERSION_PKG).version=$(VERSION) -X $(VERSION_PKG).commit=$(COMMIT) -X $(VERSION_PKG).date=$(BUILD_DATE)


# Default target
all: build

# Install development tools
setup:
	@echo "Installing development tools..."
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "✅ Development tools installed"

# Build the application
build:
	@echo "Building tabstack..."
	@go build -ldflags "$(LDFLAGS)" -o tabstack .
	@echo "✅ Built: tabstack"

# Run the application
run: build
	./tabstack

# Clean build artifacts
clean:
	@rm -f tabstack
	@rm -f *.db
	@echo "✅ Cleaned"

# Lint code
lint:
	@test -f $(HOME)/go/bin/golangci-lint || { \
		echo "❌ golangci-lint not found. Install with: make setup"; \
		exit 1; \
	}
	@echo "Running linters..."
	@$(HOME)/go/bin/golangci-lint run --timeout 5m
	@echo "✅ Lint complete"

# Format code (skips vendored dependencies)
format:
	@go fmt ./...
	@test -f $(HOME)/go/bin/gofumpt || { \
		echo "❌ gofumpt not found. Install with: make setup"; \
		exit 1; \
	}
	@$(HOME)/go/bin/gofumpt -l -w cmd internal main.go
	@echo "✅ Format complete"

# Run tests
test:
	@echo "Running tests..."
	@go test ./...
	@echo "✅ Tests complete"

# Local patches we layer over the vendored SDK (see vendor/.../local_patch_*.go).
# `go mod vendor` strips these out because they're not in modules.txt; we
# restore them from git after vendoring.
SDK_PATCHES := vendor/github.com/stainless-sdks/tabstack-go/local_patch_file_id_unions.go

# Refresh vendored dependencies after a go.mod change
vendor:
	@echo "Vendoring dependencies..."
	@go mod vendor
	@echo "Restoring local SDK patches..."
	@git restore --source HEAD -- $(SDK_PATCHES)
	@echo "✅ vendor/ refreshed (with local patches restored)"

# Verify vendor/ is in sync with go.mod (for CI)
vendor-check:
	@echo "Checking vendor/ is in sync with go.mod..."
	@go mod vendor
	@git diff --quiet vendor/ go.mod go.sum || { \
		echo "❌ vendor/ is out of sync with go.mod. Run 'make vendor' and commit the result."; \
		git diff --stat vendor/ go.mod go.sum; \
		echo "--- diff ---"; \
		git diff vendor/ go.mod go.sum | head -60; \
		exit 1; \
	}
	@echo "✅ vendor/ is in sync"
