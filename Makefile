PWD := $(shell pwd)

compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(PWD)/build/api .

compress:
	@upx $(PWD)/build/api

build: compile compress

dev:
	@air -c .air.linux.conf -d

.PHONY: compile
