.PHONY: build run clean test deploy-prep help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=portfolio-tui
BINARY_UNIX=$(BINARY_NAME)_unix

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) -o $(BINARY_NAME) -v

# Run the application
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

# Run without building (assumes already built)
run-only:
	./$(BINARY_NAME)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Build for Linux (for deployment)
build-linux:
	@echo "Building for Linux..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

# Prepare for deployment (creates deployment directory)
deploy-prep: build-linux
	@echo "Preparing deployment package..."
	mkdir -p deployment
	cp $(BINARY_UNIX) deployment/
	cp -r scripts deployment/ 2>/dev/null || true
	@echo "Deployment package ready in ./deployment/"

# Development mode with live reload (requires entr)
dev:
	@echo "Starting development mode..."
	@echo "Install entr first: brew install entr (macOS) or apt-get install entr (Linux)"
	find . -name '*.go' | entr -c make run

# Format code
fmt:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

# Run linter (requires golangci-lint)
lint:
	@echo "Running linter..."
	@echo "Install golangci-lint first: brew install golangci-lint"
	golangci-lint run

# Show help
help:
	@echo "Available targets:"
	@echo "  make build       - Build the application"
	@echo "  make run         - Build and run the application"
	@echo "  make run-only    - Run without rebuilding"
	@echo "  make clean       - Remove build artifacts"
	@echo "  make test        - Run tests"
	@echo "  make deps        - Download and tidy dependencies"
	@echo "  make build-linux - Build for Linux deployment"
	@echo "  make deploy-prep - Prepare deployment package"
	@echo "  make dev         - Run in development mode with auto-reload"
	@echo "  make fmt         - Format Go code"
	@echo "  make lint        - Run linter"
	@echo "  make help        - Show this help message"

# Default target
all: build