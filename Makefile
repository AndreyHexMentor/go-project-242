
BINARY_NAME=hexlet-path-size
BINARY_PATH=bin/$(BINARY_NAME)
SRC_PATH=./cmd/hexlet-path-size

.PHONY: all build run lint fmt test clean run_human setup-testdata

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
	go test -v ./...

setup-testdata:
	@mkdir -p testdata/dir/nested
	@echo "hello" > testdata/file.txt
	@echo "a" > testdata/dir/a.txt
	@echo "hidden" > testdata/dir/.hidden
	@echo "nested" > testdata/dir/nested/nested.txt
	@echo "Привет 🌍" > testdata/юникод.txt
	@if [ ! -e testdata/symlink ]; then ln -s file.txt testdata/symlink || true; fi
	@echo "testdata prepared"

clean:
	rm -f $(BINARY_PATH)