
BIN_DIR=./bin/
MAIN_PACKAGE=./cmd/watts-up/

# Default target
.DEFAULT_GOAL=all

# Target to clean bin directory
clean:
	@echo "Cleaning bin directory..."
	@rm -rf $(BIN_DIR)/*

# Target to build the program
build:
	@echo "Building program..."
	@go build -o $(BIN_DIR)/watts $(MAIN_PACKAGE)

# Target to clean and build the program
all: clean build
