# dock-docs Makefile
# ──────────────────────────────────────────────────────────

MODULE   := github.com/northcutted/dock-docs
BINARY   := bin/dock-docs
COVFILE  := coverage.out
# Packages to include in coverage (excludes main.go root package which
# is just a 3-line entry point calling cmd.Execute()).
COVPKGS  := ./cmd/...,./pkg/...

VERSION  ?= dev
COMMIT   ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE     ?= $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

LDFLAGS  := -s -w \
	-X $(MODULE)/cmd.Version=$(VERSION) \
	-X $(MODULE)/cmd.Commit=$(COMMIT) \
	-X $(MODULE)/cmd.Date=$(DATE)

.PHONY: build test lint lint-fix fmt vet tidy cover cover-html clean ci help

## help: show this help message
help:
	@echo "Usage: make <target>"
	@echo ""
	@sed -n 's/^## //p' $(MAKEFILE_LIST) | column -t -s ':'

## build: compile the binary (with version info)
build:
	go build -ldflags '$(LDFLAGS)' -o $(BINARY) .

## test: run all tests with race detection
test:
	go test -race -v ./...

## lint: run golangci-lint
lint:
	golangci-lint run

## lint-fix: run golangci-lint with auto-fix
lint-fix:
	golangci-lint run --fix

## fmt: format code
fmt:
	go fmt ./...

## vet: run go vet
vet:
	go vet ./...

## tidy: tidy and verify go.mod
tidy:
	go mod tidy
	go mod verify

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

## ci: run the full CI pipeline locally (fmt, vet, lint, test)
ci: fmt vet lint test

## clean: remove build artifacts
clean:
	rm -f $(BINARY) $(COVFILE)
