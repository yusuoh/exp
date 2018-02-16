# Simple makefile.
#
all: deps
	go build ./...

.PHONY: deps

deps:
	go get -v -t -d ./...
