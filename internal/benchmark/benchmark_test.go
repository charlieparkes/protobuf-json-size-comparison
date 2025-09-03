package benchmark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonProtoLengthsSize1(t *testing.T) {
	policy := CreatePolicyData(1)
	assert.NotNil(t, policy, "CreatePolicyData should not return nil")

	j, gj, p, gp := JsonProtoLengths(policy)
	assert.Greater(t, j, 0, "JSON length should be positive")
	assert.Greater(t, gj, 0, "Gzipped JSON length should be positive")
	assert.Greater(t, p, 0, "Protobuf length should be positive")
	assert.Greater(t, gp, 0, "Gzipped protobuf length should be positive")

	t.Logf("Sizes - JSON: %d, Gzipped JSON: %d, Proto: %d, Gzipped Proto: %d", j, gj, p, gp)
}

func TestJsonProtoLengthsSize100(t *testing.T) {
	policy := CreatePolicyData(100)
	assert.NotNil(t, policy, "CreatePolicyData should not return nil")

	j, gj, p, gp := JsonProtoLengths(policy)
	assert.Greater(t, j, 0, "JSON length should be positive")
	assert.Greater(t, gj, 0, "Gzipped JSON length should be positive")
	assert.Greater(t, p, 0, "Protobuf length should be positive")
	assert.Greater(t, gp, 0, "Gzipped protobuf length should be positive")

	t.Logf("Sizes - JSON: %d, Gzipped JSON: %d, Proto: %d, Gzipped Proto: %d", j, gj, p, gp)
}
