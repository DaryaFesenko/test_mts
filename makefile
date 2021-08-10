APP = test_mts
PROJECT = github.com/DaryaFesenko/test_mts
STDERR=/tmp/.$(APP)-stderr.txt

HAS_LINT := $(shell command -v golangci-lint;)
HAS_IMPORTS := $(shell command -v goimports;)

all: run

lint: bootstrap
	@echo "+ $@"
	@golangci-lint run

run: clean build
	@echo "+ $@"
	./${APP}

build: lint
	@echo "+ $@"
	@go build

clean:
	@echo "+ $@"
	@rm -f ./${APP}

bootstrap:
	@echo "+ $@"
ifndef HAS_LINT
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1
endif
ifndef HAS_IMPORTS
	go get -u golang.org/x/tools/cmd/goimports
endif
	

.PHONY: all \
	lint \
	run \
	build \
	clean \
	bootstrap