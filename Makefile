# Setup package name variables
NAME := ytd
PKG := github.com/ch3ck/$(NAME)
PREFIX?=$(shell pwd)
BUILDTAGS=

.PHONY: clean all fmt vet build test install static
.DEFAULT: default

all: clean fmt vet build test install

build: fmt vet
	@echo "+ $@"
	@go build -tags "$(BUILDTAGS) cgo" .

static:
	@echo "+ $@"
	CGO_ENABLED=1 go build -tags "$(BUILDTAGS) cgo static_build" -ldflags "-w -extldflags -static" -o ytd .

fmt:
	@echo "+ $@"
	@gofmt -s -l -w . | grep -v vendor | tee /dev/stderr

test:
	@echo "+ $@"
	@find . -name \*.mp3 -delete #clean previous test files.
	@go test -v -tags "$(BUILDTAGS) cgo" $(shell go list ./... | grep -v vendor)
	@find . -name \*.mp3 -delete # clean previous test downloads
	@go test -bench=. $(shell go list ./... | grep -v vendor)

vet:
	@echo "+ $@"
	@go vet $(shell go list ./... | grep -v vendor)

clean:
	@echo "+ $@"
	@rm -rf ytd
	@find . -name \*.mp3 -delete
	@find . -name \*.flv -delete

install:
	@echo "+ $@"
	@docker build -t ch3ck/ytd:v1 . 
	@go install .
