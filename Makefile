PROJECT_NAME:=protobuf-json-size-comparison
PROTO_OUT=internal/pb
include shared.mk

.PHONY: run
run:
	go run cmd/comparison/main.go
