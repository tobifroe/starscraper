.PHONY: build build-alpine clean test help default



BIN_NAME=starscraper

VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')

default: test

help:
	@echo 'Management commands for starscraper:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs go mod download, mostly used for ci.'
	@echo '    make build-alpine    Compile optimized for alpine linux.'
	@echo '    make clean           Clean the directory tree.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/tobifroe/starscraper/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/tobifroe/starscraper/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

get-deps:
	go mod download

build-alpine:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '-s -w -linkmode external -extldflags "-static" -X github.com/tobifroe/starscraper/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/tobifroe/starscraper/version.BuildDate=${BUILD_DATE}' -o bin/${BIN_NAME} && upx --brute bin/${BIN_NAME} && tar -czvf bin/${BIN_NAME}-${VERSION}.tar.gz bin/${BIN_NAME} && rm bin/${BIN_NAME}

build-alpine-mac:
	@echo "building ${BIN_NAME} ${VERSION} for Mac"
	@echo "GOPATH=${GOPATH}"
	@echo "GOOS=darwin"
	@echo "GOARCH=amd64"
	env GOOS=darwin GOARCH=amd64 GOHOSTARCH=amd64 GOHOSTOS=linux go build -ldflags '-w -X github.com/tobifroe/starscraper/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/tobifroe/starscraper/version.BuildDate=${BUILD_DATE}' -o bin/${BIN_NAME} && upx --brute bin/${BIN_NAME} && tar -czvf bin/${BIN_NAME}-macos-${VERSION}.tar.gz bin/${BIN_NAME} && rm bin/${BIN_NAME}

test:
	go run gotest.tools/gotestsum@latest