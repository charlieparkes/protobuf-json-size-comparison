package main

import (
	"fmt"

	"github.com/charlieparkes/protobuf-json-size-comparison/internal/benchmark"
)

func main() {
	fmt.Printf("|%15s | %10s | %15s | %10s | %15s | %21s | %40s | \n", "no of policies", "json", " gzipped json", " proto", "gzipped proto", "proto size(%) of json", "gzipped proto size(%) of gzipped json")
	for _, dataSize := range []int{500000, 1000000, 5000000, 10000000} {
		protoStruct := benchmark.CreatePolicyData(dataSize)
		jsonl, gzJsonlen, protol, gzProto := benchmark.JsonProtoLengths(protoStruct)
		fmt.Printf("|%15d | %10d | %15d | %10d | %15d | %21f | %40f | \n", dataSize, jsonl, gzJsonlen, protol, gzProto, float32(protol)/float32(jsonl), float32(gzProto)/float32(gzJsonlen))
	}
}
