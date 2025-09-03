package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	examplev1 "github.com/charlieparkes/protobuf-json-size-comparison/internal/pb/example/v1"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Printf("|%15s | %10s | %15s | %10s | %15s | %21s | %40s | \n", "no of tickers", "json", " gzipped json", " proto", "gzipped proto", "proto size(%) of json", "gzipped proto size(%) of gzipped json")
	for _, dataSize := range []int{500000, 1000000, 5000000, 10000000} {
		protoStruct := createTestDatata(dataSize)
		jsonl, gzJsonlen, protol, gzProto := jsonProtoLengts(protoStruct)
		fmt.Printf("|%15d | %10d | %15d | %10d | %15d | %21f | %40f | \n", dataSize, jsonl, gzJsonlen, protol, gzProto, float32(protol)/float32(jsonl), float32(gzProto)/float32(gzJsonlen))
	}

}
func createTestDatata(numberOfEntries int) *examplev1.Test {

	tickers := make([]*examplev1.Ticker, numberOfEntries)

	for i := 0; i < numberOfEntries; i++ {
		tickers[i] = &examplev1.Ticker{
			Value: rand.Float32() * 10.0,
			Name:  RandStringRunes(3),
		}
	}

	return &examplev1.Test{
		Query:         "myQuery",
		PageNumber:    42,
		ResultPerPage: 100,
		Tickers:       tickers,
	}
}

func jsonProtoLengts(protoSome *examplev1.Test) (jsonLen, gzipJSONLen, protoLen, gzpProto int) {
	data, _ := proto.Marshal(protoSome)
	protoLen = len(data)
	jsonified, _ := json.Marshal(protoSome)
	jsonLen = len(jsonified)
	gzipJSONLen = gzipLen(jsonified)
	gzpProto = gzipLen(data)
	return
}

func gzipLen(jsonData []byte) int {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(jsonData); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}

	return b.Len()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rnd.Intn(len(letterRunes))]
	}
	return string(b)
}
