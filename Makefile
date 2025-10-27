
BINARY_NAME=hexlet-path-size
BINARY_PATH=bin/$(BINARY_NAME)
SRC_PATH=./cmd/hexlet-path-size

.PHONY: all build run lint fmt test clean run_human

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(dir $(BINARY_PATH))
	go build -o $(BINARY_PATH) $(GOFLAGS) $(SRC_PATH)
	@echo "$(BINARY_NAME) built successfully to $(BINARY_PATH)"

run: build
	@echo "Running $(BINARY_NAME) with default path $(DEFAULT_PATH) and ARGS..."
	$(BINARY_PATH) $(DEFAULT_PATH) $(ARGS)

run_human: build
	@echo "Running $(BINARY_NAME) with --human..."
	$(BINARY_PATH) $(DEFAULT_PATH)

lint:
	@echo "Running linters..."
	golangci-lint run ./...

fmt:
	@echo "Formatting code..."
	goimports -w .

test:
	@go mod tidy
	@echo "Cleaning testdata..."
	@rm -rf ./testdata
	@echo "Running tests..."
	go test -v ./code/...

clean:
	rm -f $(BINARY_PATH)