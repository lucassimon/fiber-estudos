PWD := $(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(PWD)/build/api $(PWD)/src/

compress:
	@upx $(PWD)/build/api

.production: build compress

.dev: build compress

.PHONY: build compress
