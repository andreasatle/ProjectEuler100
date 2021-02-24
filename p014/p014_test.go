package p014_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p014"
	"github.com/stretchr/testify/assert"
)

func TestMaxCollatzSequenceV1(t *testing.T) {
	assert.Equal(t, 837799, p014.MaxCollatzSequenceV1(1000000))
}

func TestMaxCollatzSequenceV2(t *testing.T) {
	assert.Equal(t, 837799, p014.MaxCollatzSequenceV2(1000000))
}

func TestMaxCollatzSequenceV3(t *testing.T) {
	assert.Equal(t, 837799, p014.MaxCollatzSequenceV3(1000000))
}

func TestMaxCollatzSequenceV4(t *testing.T) {
	assert.Equal(t, 837799, p014.MaxCollatzSequenceV4(1000000))
}

func BenchmarkMaxCollatzSequenceV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p014.MaxCollatzSequenceV1(1000000)
	}
}

func BenchmarkMaxCollatzSequenceV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p014.MaxCollatzSequenceV2(1000000)
	}
}

func BenchmarkMaxCollatzSequenceV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p014.MaxCollatzSequenceV3(1000000)
	}
}

func BenchmarkMaxCollatzSequenceV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p014.MaxCollatzSequenceV4(1000000)
	}
}
