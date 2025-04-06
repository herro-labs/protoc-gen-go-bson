PLUGIN_NAME=protoc-gen-go-bson
PLUGIN_BIN=./plugins/$(PLUGIN_NAME)

export PATH := $(abspath $(dir $(PLUGIN_BIN))):$(PATH)

.PHONY: build

build:
	go build -o $(PLUGIN_BIN) main.go
