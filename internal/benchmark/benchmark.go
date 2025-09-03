package benchmark

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"math/rand"
	"time"

	examplev1 "github.com/charlieparkes/protobuf-json-size-comparison/internal/pb/example/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// CreatePolicyData generates test policy data with the specified number of coverage entries
func CreatePolicyData(numberOfEntries int) *examplev1.Policy {
	drivers := []*examplev1.Driver{
		{
			CustomerId: "driver-123",
			Role:       examplev1.DriverRole_DRIVER_ROLE_PRIMARY,
		},
		{
			CustomerId: "driver-456",
			Role:       examplev1.DriverRole_DRIVER_ROLE_SECONDARY,
		},
	}

	vehicles := []*examplev1.Vehicle{
		{
			Vin:           "1HGBH41JXMN109186",
			Make:          "Honda",
			Model:         "Civic",
			Year:          2021,
			YearlyMileage: 12000,
		},
	}

	coverages := make([]*examplev1.CoverageEntry, numberOfEntries)

	for i := 0; i < numberOfEntries; i++ {
		coverages[i] = &examplev1.CoverageEntry{
			CoverageId:  fmt.Sprintf("COV%d", i),
			OptionId:    fmt.Sprintf("OPT%d", i),
			LimitAmount: rand.Float32() * 1000000.0,
		}
	}

	return &examplev1.Policy{
		PolicyNumber:   "POL-2024-001",
		RevisionNumber: 1,
		CustomerId:     "customer-789",
		Drivers:        drivers,
		StateCode:      "CA",
		Vehicles:       vehicles,
		Premium:        2450.50,
		AgentId:        "agent-abc",
		Coverages:      coverages,
	}
}

// JsonProtoLengths calculates the serialized sizes for JSON and protobuf formats
func JsonProtoLengths(protoSome *examplev1.Policy) (jsonLen, gzipJSONLen, protoLen, gzipProtoLen int) {
	data, _ := proto.Marshal(protoSome)
	protoLen = len(data)
	jsonified, _ := protojson.Marshal(protoSome)
	jsonLen = len(jsonified)
	gzipJSONLen = gzipLen(jsonified)
	gzipProtoLen = gzipLen(data)
	return jsonLen, gzipJSONLen, protoLen, gzipProtoLen
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

// RandStringRunes generates a random string of the specified length
func RandStringRunes(n int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rnd.Intn(len(letterRunes))]
	}
	return string(b)
}
