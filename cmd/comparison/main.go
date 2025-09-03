package main

import (
	"fmt"

	"github.com/charlieparkes/protobuf-json-size-comparison/internal/benchmark"
)

func main() {
	fmt.Printf("|%15s | %10s | %15s | %10s | %15s | %21s | %21s | %40s | \n", "no of elements", "json", " gzipped json", " proto", "gzipped proto", "proto size(%) of json", "gzipped proto(%) of json", "gzipped proto size(%) of gzipped json")
	for _, dataSize := range []int{1, 2, 3, 4, 5, 10, 100, 1000, 10000} {
		protoStruct := benchmark.CreatePolicyData(dataSize)
		jsonl, gzJsonlen, protol, gzProto := benchmark.JsonProtoLengths(protoStruct)
		fmt.Printf("|%15d | %10d | %15d | %10d | %15d | %20.0f%% | %23.0f%% | %39.0f%% | \n", dataSize, jsonl, gzJsonlen, protol, gzProto, float32(protol)/float32(jsonl)*100, float32(gzProto)/float32(jsonl)*100, float32(gzProto)/float32(gzJsonlen)*100)
	}
}
