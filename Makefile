APP_NAME := devisor-api
BIN := bin/$(APP_NAME)
GO := go
GOBIN := $(BIN)

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	$(GO) build -o $(GOBIN) ./cmd/server/main.go

# Run the application
.PHONY: run
run:
	@echo "Running $(APP_NAME)..."
	$(GO) run ./cmd/server/main.go

# Clean binaries
.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf bin/*

# Install dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	$(GO) mod tidy
	$(GO) mod download

# Test
.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test -v ./...

# Lint
.PHONY: lint
lint:
	@echo "Linting..."
	golangci-lint run

# Database migrations

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting Go code..."
	$(GO) fmt ./...