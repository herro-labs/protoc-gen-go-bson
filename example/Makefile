PROTO_DIR=proto
OUT_DIR=protogen
PROTO_FILES=$(wildcard $(PROTO_DIR)/*.proto)

generate:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-bson_out=$(OUT_DIR) \
		--go-bson_opt=paths=source_relative \
		$(PROTO_FILES)