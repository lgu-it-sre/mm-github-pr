BINARY_NAME=plugin

.PHONY: default
default: build deploy

.PHONY: build
build:
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME) -v
	@tar -czf bin/$(BINARY_NAME).tar.gz plugin.json assets/icon.png -C bin/ $(BINARY_NAME)

.PHONY: deploy
deploy:
	@bin/mmctl plugin delete com.github.lgu-it-sre.mm-github-pr
	@bin/mmctl plugin add bin/$(BINARY_NAME).tar.gz
	@bin/mmctl plugin enable com.github.lgu-it-sre.mm-github-pr

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf bin/
