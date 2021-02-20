package p001_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p001"
	"github.com/stretchr/testify/assert"
)

func TestCountV1(t *testing.T) {
	assert.Equal(t, 233168, p001.CountV1(3, 5, 1000))
}

func TestCountV2(t *testing.T) {
	assert.Equal(t, 233168, p001.CountV2(3, 5, 1000))
}

func TestCountV3(t *testing.T) {
	assert.Equal(t, 233168, p001.CountV3(3, 5, 1000))
}

func BenchmarkCountV1N1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p001.CountV1(3, 5, 1000)
	}
}

func BenchmarkCountV2N1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p001.CountV2(3, 5, 1000)
	}
}

func BenchmarkCountV3N1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p001.CountV3(3, 5, 1000)
	}
}

func BenchmarkCountV1N100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p001.CountV1(3, 5, 100000)
	}
}

func BenchmarkCountV2N100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p001.CountV2(3, 5, 100000)
	}
}

func BenchmarkCountV3N100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p001.CountV3(3, 5, 100000)
	}
}
