# Root Directories
BINARY_ROOT_DIR := bin
BUILD_LOG_DIR := $(BINARY_ROOT_DIR)/logs

BINARY_NAME := invoices
BINARY_PATH := $(BINARY_ROOT_DIR)/$(BINARY_NAME)

all: clean proto build run

proto:
	@echo "Compiling Proto file..."
	@protoc --go_out=. protobuf/schema.proto

dir:
	@mkdir -p $(BINARY_ROOT_DIR)
	@mkdir -p $(BUILD_LOG_DIR)

build: dir
	@echo "Building Executable..."
	@go build -v -gcflags="all=-N -l" -x -ldflags="-s -w" -o $(BINARY_PATH) cmd/$(BINARY_NAME)/*.go > $(BUILD_LOG_DIR)/build.log 2>&1

run:
	@echo "Running..."
	@$(BINARY_ROOT_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning all build artifacts..."
	@rm -rf $(BINARY_ROOT_DIR)
	@rm -rf $(BUILD_LOG_DIR)
	@echo "Clean complete."

code-export:
	@git archive --format=tar.gz --output=invoice-archive-import-export.tar.gz HEAD

# Help command to display available targets
help:
	@echo "\n"
	@echo "Makefile commands:"
	@echo "  make build         - Builds the project for build"
	@echo "  make run       	- Runs the build build"
	@echo "  make clean         - Cleans all build artifacts"
	@echo "  make help          - Displays this help message"
	@echo "\n"

.PHONY: all linux mac build run-linux run-mac run clean help
