# Variables
APP_NAME := auth-service
GO_FILES := $(shell find . -name '*.go' -not -path "./vendor/*")

# Tools
FMT := gofumpt
LINTER := golangci-lint
GOOSE := goose

# Database Variables
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= user
DB_PASSWORD ?= password
DB_NAME ?= auth_db
DB_SSLMODE ?= disable

# Commands
.PHONY: all build run fmt lint test clean install-tools migrate-up migrate-down

all: install-tools fmt lint build

# Build the application
build:
	@echo "Building the application..."
	@go build -o $(APP_NAME)

# Run the application
run: build
	@echo "Running the application..."
	@./$(APP_NAME)

# Format the code
fmt: ensure-fmt
	@echo "Formatting Go files..."
	@$(FMT) -w $(GO_FILES)

# Lint the code
lint: ensure-linter
	@echo "Linting the code..."
	@$(LINTER) run

# Test the code
test:
	@echo "Running tests..."
	@go test ./... -v

# Clean up
clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Ensure gofumpt is installed
ensure-fmt:
	@command -v $(FMT) >/dev/null 2>&1 || { \
		echo "Installing gofumpt..."; \
		go install mvdan.cc/gofumpt@latest; \
	}

# Ensure golangci-lint is installed
ensure-linter:
	@command -v $(LINTER) >/dev/null 2>&1 || { \
		echo "Installing golangci-lint..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	}

# Ensure goose is installed
ensure-goose:
	@command -v $(GOOSE) >/dev/null 2>&1 || { \
		echo "Installing goose..."; \
		go install github.com/pressly/goose/v3/cmd/goose@latest; \
	}

# Install all required tools
install-tools: ensure-fmt ensure-linter ensure-goose

# Migrate up
migrate-up: ensure-goose
	@echo "Running database migrations (up)..."
	@$(GOOSE) -dir migrations postgres "$(DB_CONNECTION)" up

# Migrate down
migrate-down: ensure-goose
	@echo "Reverting database migrations (down)..."
	@$(GOOSE) -dir migrations postgres "$(DB_CONNECTION)" down

# Generate DB connection string
DB_CONNECTION := host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) sslmode=$(DB_SSLMODE)

# Initialize the project
init:
	@echo "Initializing the project..."
	@go mod init $(APP_NAME)
	@make deps
	@make install-tools

# Help command
help:
	@echo "Available make commands:"
	@echo "  make all            - Run fmt, lint, and build"
	@echo "  make build          - Build the application"
	@echo "  make run            - Build and run the application"
	@echo "  make fmt            - Format the code using gofumpt"
	@echo "  make lint           - Lint the code using golangci-lint"
	@echo "  make test           - Run tests"
	@echo "  make clean          - Clean up build artifacts"
	@echo "  make deps           - Install project dependencies"
	@echo "  make install-tools  - Install required tools (gofumpt, golangci-lint, goose)"
	@echo "  make migrate-up     - Run database migrations (up)"
	@echo "  make migrate-down   - Revert database migrations (down)"
	@echo "  make init           - Initialize the project and install dependencies"
