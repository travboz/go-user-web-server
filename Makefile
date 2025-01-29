include .env # read from .env file

OUTPUT_BINARY = diy-server
OUTPUT_DIR = ./bin
ENTRY_DIR = ./

.PHONY: build
build:
	@mkdir -p $(OUTPUT_DIR)
	go build -o $(OUTPUT_DIR)/$(OUTPUT_BINARY) $(ENTRY_DIR)

.PHONY: run
run: build
	@$(OUTPUT_DIR)/$(OUTPUT_BINARY)

.PHONY: clean
clean:
	@rm -rf $(OUTPUT_DIR)