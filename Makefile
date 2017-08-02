# Builds and runs the micro services go generate && go build
#$ bin/%: CGO_ENABLED=0 go build -ldflags '-s -w' -tags netgo -v -o $@ ./cmd/$*
# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)
BUILDTAGS=

.PHONY: clean all fmt vet lint build test install static
.DEFAULT: default

all: clean build fmt lint test vet install

build:
	@echo "+ $@"
	@sh ./build.sh
	@go build -tags "$(BUILDTAGS) cgo" . ## Thinking of downloading the required go deps in here.
	#@docker build -t ch3ck/ytd:v1 .

static:
	@echo "+ $@"
	CGO_ENABLED=1 go build -tags "$(BUILDTAGS) cgo static_build" -ldflags "-w -extldflags -static" -o ytd .

fmt:
	@echo "+ $@"
	@gofmt -s -l . | grep -v vendor | tee /dev/stderr

lint:
	@echo "+ $@"
	@golint ./... | grep -v vendor | tee /dev/stderr

test: fmt lint vet
	@echo "+ $@"
	@go test -v -tags "$(BUILDTAGS) cgo" $(shell go list ./... | grep -v vendor)

vet:
	@echo "+ $@"
	@go vet $(shell go list ./... | grep -v vendor)

clean:
	@echo "+ $@"
	@rm -rf ytd

install:
	@echo "+ $@"
	@go install .
