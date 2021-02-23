package p010_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p010"
	"github.com/stretchr/testify/assert"
)

func TestPrimeSum(t *testing.T) {
	assert.Equal(t, 17, p010.PrimeSumV1(10))
	assert.Equal(t, 17, p010.PrimeSumV2(10))
	assert.Equal(t, 142913828922, p010.PrimeSumV1(2000000))
	assert.Equal(t, 142913828922, p010.PrimeSumV2(2000000))
}

func BenchmarkPrimeSumV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p010.PrimeSumV1(2000000)
	}
}

func BenchmarkPrimeSumV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p010.PrimeSumV2(2000000)
	}
}
