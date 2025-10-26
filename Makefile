
BINARY_NAME=hexlet-path-size
BINARY_PATH=bin/$(BINARY_NAME)
SRC_PATH=./cmd/hexlet-path-size

build:
	go build -o $(BINARY_PATH) $(SRC_PATH)

run: build
	$(BINARY_PATH) $(ARGS)

lint:
	golangci-lint run ./...

fmt:
	goimports -w .

test:
	go mod tidy
	go test -v ./...

install:
	go install

clean:
	rm -f $(BINARY_PATH)