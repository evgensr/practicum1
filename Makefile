GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
GOMOD=$(GOCMD) mod
GOPATH?=`$(GOCMD) env GOPATH`

BINARY=shortener
BINARY_LINUX=$(BINARY)_linux

TESTS=./...
COVERAGE_FILE=coverage.out

BUILD=build/


BUILD_VERSION=$(shell git tag|tail -n 1)
BUILD_NUMBER=$(strip $(if $(TRAVIS_BUILD_NUMBER), $(TRAVIS_BUILD_NUMBER), 0))
BUILD_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)


all: download test build
.PHONY: all download test build build-linux install build-staticlint test-bench

download:
	@echo "[*] $@"
	@$(GOMOD) download

test: download
	@echo "[*] $@"
	$(GOTEST) -v $(TESTS)

test-bench: download
	@echo "[*] $@"
	$(GOTEST) -v $(TESTS) -bench . -benchmem


coverage:
	@echo "[*] $@"
	$(GOTEST) -coverprofile=$(COVERAGE_FILE) $(TESTS)
	$(GOTOOL) cover -html=$(COVERAGE_FILE)

build: download
	@echo "[*] $@"
	$(GOBUILD) -o $(BINARY)   -ldflags "-X 'main.buildCommit=${BUILD_COMMIT}' -X main.buildVersion=${BUILD_VERSION}.${BUILD_NUMBER} -X main.buildDate=${BUILD_DATE} "  -v ./cmd/shortener

build-staticlint: download
	@echo "[*] $@"
	$(GOBUILD) -o staticlint -v ./cmd/staticlint

build-linux: download
	@echo "[*] $@"
	GOOS="linux" GOARCH="amd64" $(GOBUILD) -o $(BINARY_LINUX) -v -ldflags "-X 'main.buildCommit=${BUILD_COMMIT}'" -v ./cmd/shortener

