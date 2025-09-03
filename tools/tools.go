//go:build tools
// +build tools

// Ensure that tool versions are tracked in go mod.
// https://github.com/golang/go/wiki/Modules#how-can-i-grack-tool-dependencies-for-a-module
package tools

import (
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
