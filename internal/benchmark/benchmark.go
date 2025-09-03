package benchmark

import (
	"bytes"
	"compress/gzip"
	"math/rand"
	"time"

	examplev1 "github.com/charlieparkes/protobuf-json-size-comparison/internal/pb/example/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// CreatePolicyData generates test policy data with the specified number of entries for each collection
func CreatePolicyData(numberOfEntries int) *examplev1.Policy {
	drivers := make([]*examplev1.Driver, numberOfEntries)
	for i := 0; i < numberOfEntries; i++ {
		role := examplev1.DriverRole_DRIVER_ROLE_PRIMARY
		if i%3 == 1 {
			role = examplev1.DriverRole_DRIVER_ROLE_SECONDARY
		} else if i%3 == 2 {
			role = examplev1.DriverRole_DRIVER_ROLE_LISTED
		}

		drivers[i] = &examplev1.Driver{
			CustomerId: randString(12), // Generate 12-character customer ID
			Role:       role,
		}
	}

	vehicles := make([]*examplev1.Vehicle, numberOfEntries)
	for i := 0; i < numberOfEntries; i++ {
		vehicles[i] = &examplev1.Vehicle{
			Vin:           randString(17),                 // VIN is typically 17 characters
			Make:          randString(8),                  // Car make (e.g., "Honda", "Toyota")
			Model:         randString(10),                 // Car model (e.g., "Civic", "Camry")
			Year:          int64(2015 + rand.Intn(10)),    // Random year between 2015-2024
			YearlyMileage: int64(5000 + rand.Intn(20000)), // Random mileage 5k-25k
		}
	}

	coverages := make([]*examplev1.CoverageEntry, numberOfEntries)
	for i := 0; i < numberOfEntries; i++ {
		coverages[i] = &examplev1.CoverageEntry{
			CoverageId:  randString(6),              // 6-character coverage ID
			OptionId:    randString(8),              // 8-character option ID
			LimitAmount: rand.Float32() * 1000000.0, // Random limit up to $1M
		}
	}

	return &examplev1.Policy{
		PolicyNumber:   randString(15),           // 15-character policy number
		RevisionNumber: int32(1 + rand.Intn(10)), // Random revision 1-10
		CustomerId:     randString(12),           // 12-character customer ID
		Drivers:        drivers,
		StateCode:      randString(2), // 2-character state code
		Vehicles:       vehicles,
		Premium:        rand.Float32() * 5000.0, // Random premium up to $5000
		AgentId:        randString(10),          // 10-character agent ID
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

// randString generates a random string of the specified length
func randString(n int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rnd.Intn(len(letterRunes))]
	}
	return string(b)
}
