# dock-docs Makefile
# ──────────────────────────────────────────────────────────

BINARY   := bin/dock-docs
COVFILE  := coverage.out
# Packages to include in coverage (excludes main.go root package which
# is just a 3-line entry point calling cmd.Execute()).
COVPKGS  := ./cmd/...,./pkg/...

.PHONY: build test lint fmt cover cover-html clean

## build: compile the binary
build:
	go build -o $(BINARY) .

## test: run all tests with race detection
test:
	go test -race -v ./...

## lint: run golangci-lint
lint:
	golangci-lint run

## fmt: format code
fmt:
	go fmt ./...

## cover: run tests with coverage, excluding the root package (main.go)
cover:
	go test -coverprofile=$(COVFILE) -coverpkg=$(COVPKGS) ./cmd/... ./pkg/...
	@echo ""
	@echo "── Per-function coverage ──"
	@go tool cover -func=$(COVFILE)
	@echo ""
	@echo "── Summary ──"
	@go tool cover -func=$(COVFILE) | tail -1

## cover-html: open coverage report in browser
cover-html: cover
	go tool cover -html=$(COVFILE)

## clean: remove build artifacts
clean:
	rm -f $(BINARY) $(COVFILE)
