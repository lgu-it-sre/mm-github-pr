BINARY_NAME=plugin

.PHONY: default
default: build

.PHONY: build
build:
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME) -v
	@tar -czf bin/$(BINARY_NAME).tar.gz plugin.json -C bin/ $(BINARY_NAME)

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf bin/
