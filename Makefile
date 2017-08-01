SHELL := bash

NAME    := inspeqtor
VERSION ?= $(shell git describe --tags --abbrev=0)

# Wrapped package path
ORIGPKG := $(GOPATH)/src/github.com/mperham/inspeqtor

all: test

deps:
	@echo "+ $@"
	go get -d github.com/mperham/inspeqtor/...
	go get github.com/jteeuwen/go-bindata/...

bin: deps fmt generate
	@echo "+ $@"
	pushd $(ORIGPKG) >/dev/null && \
		go-bindata -pkg inspeqtor -o templates.go templates/... && \
	popd >/dev/null
	go build -ldflags "-X main.version=$(VERSION)" -o $(NAME) main.go

test: deps fmt generate
	@echo "+ $@"
	@go test -v ./...

generate:
	@echo "+ $@"
	@go generate

fmt:
	@echo "+ $@"
	@find . -name "*.go" -exec gofmt -l -w {} \;

clean:
	@echo "+ $@"
	@$(RM) -f $(NAME)

.PHONY: all clean deps bin fmt generate test
