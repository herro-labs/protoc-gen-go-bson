PLUGIN_NAME=protoc-gen-go-bson
PROTO_DIR=proto
OUT_DIR=pb
PLUGIN_BIN=./plugins/$(PLUGIN_NAME)
PROTO_FILES=$(wildcard $(PROTO_DIR)/*.proto)

export PATH := $(abspath $(dir $(PLUGIN_BIN))):$(PATH)

.PHONY: all build generate

all: build generate

generate:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-bson_out=$(OUT_DIR) \
		--go-bson_opt=paths=source_relative \
		$(PROTO_FILES)

build:
	go build -o $(PLUGIN_BIN) main.go
