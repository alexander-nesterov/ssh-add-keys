SHELL := /bin/sh

APP_NAME := ssh-add-keys

.PHONY: init
init:
	go mod init $(APP_NAME)

.PHONY: get-dep
get-dep:
	go get .

.PHONY: build
build:
	go build -o bin/$(APP_NAME)

.PHONY: run
run:
	go run .