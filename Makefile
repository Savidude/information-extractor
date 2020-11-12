# ------------------------------------------------------------------------
#
# Copyright (c) 2020, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
#
# ------------------------------------------------------------------------

PROJECT_ROOT := $(realpath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
BUILD_DIRECTORY := build
BUILD_ROOT := $(PROJECT_ROOT)/$(BUILD_DIRECTORY)
MAIN_PACKAGE := information-extractor
BUILD_VERSION := 1.1.1

all: clean build-linux build-darwin build-windows

.PHONY: build-linux
build-linux: ## Builds a Linux executable
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH \
    -ldflags "-X main.version=$(BUILD_VERSION) -X 'main.buildDate=$$(date -u '+%Y-%m-%d %H:%M:%S UTC')'" \
	-o $(BUILD_ROOT)/target/linux/$(MAIN_PACKAGE)_linux -x $(PROJECT_ROOT)/cmd/$(MAIN_PACKAGE)


.PHONY: build-darwin
build-darwin: ## Builds a Darwin executable
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 \
	go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH \
    -ldflags "-X main.version=$(BUILD_VERSION) -X 'main.buildDate=$$(date -u '+%Y-%m-%d %H:%M:%S UTC')'" \
	-o $(BUILD_ROOT)/target/darwin/$(MAIN_PACKAGE)_darwin -x $(PROJECT_ROOT)/cmd/$(MAIN_PACKAGE)

.PHONY: build-windows
build-windows: ## Builds a Windows executable
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
	go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH \
    -ldflags "-X main.version=$(BUILD_VERSION) -X 'main.buildDate=$$(date -u '+%Y-%m-%d %H:%M:%S UTC')'" \
    -o $(BUILD_ROOT)/target/windows/$(MAIN_PACKAGE)_windows.exe -x $(PROJECT_ROOT)/cmd/$(MAIN_PACKAGE)

.PHONY: test
test: ## Runs the tests
	go test -v ./pkg/...

.PHONY: clean
clean: ## Cleans up build artifacts
	rm -fr $(BUILD_ROOT)/target

.PHONY: setup-lint
setup-lint: ## Install golint
	go get -u golang.org/x/lint/golint

.PHONY: run-lint
run-lint: ## Run golint on the code
	golint  ./pkg/* ./cmd/*

.PHONY: format
format: ## Run gofmt on the code
	gofmt -w ./pkg/* ./cmd/*

.PHONY: version
version: ## Get project build version
	@echo ${BUILD_VERSION}

.PHONY: help
help: ## Shows the help
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
        awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ''
