# defalut
.PHONY: all
all: help

# define variables
BINARY_NAME := goginmvc
VERSION := $(shell git describe --tags --always --dirty)
GOFLAGS := -ldflags "-X main.Version=$(VERSION)"
BUILD_DIR := build
SRC_DIRS := ./...

# set variables
ifeq ($(origin env), undefined)
    ENV := dev
else ifeq ($(env), dev)
    ENV := dev
    GOFLAGS += -ldflags "-X main.Env=development"
else ifeq ($(env), prod)
    ENV := prod
    GOFLAGS += -ldflags "-X main.Env=production"
endif

# help info
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Available targets:"
	@echo "  build        Build the application for development or production."
	@echo "  test         Run all unit tests."
	@echo "  clean        Remove build artifacts."
	@echo "  install      Install dependencies."
	@echo "  fmt          Format Go source files."
	@echo "  vet          Run go vet to check for issues in the code."

# build
.PHONY: build
build:
	@echo "Building $(BINARY_NAME) for environment: $(ENV)"
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .

.PHONY: Dev
dev:
	@echo "Starting development server..."
	@$(MAKE) build env=dev
	@go run $(SRC_DIRS)

# execute unit tests
.PHONY: test
test:
	@echo "Running tests..."
	go test -v $(SRC_DIRS)

# clean build
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# install dependencies
.PHONY: install
install:
	@echo "Installing dependencies..."
	go mod download

# fmt
.PHONY: fmt
fmt:
	@echo "Formatting Go source files..."
	gofmt -w $(SRC_DIRS)

# vet
.PHONY: vet
vet:
	@echo "Running go vet..."
	go vet $(SRC_DIRS)