package p012_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p012"
	"github.com/stretchr/testify/assert"
)

func TestFindTrianguleNumberWithNDivisorsV1N5(t *testing.T) {
	assert.Equal(t, 28, p012.FindTrianguleNumberWithNDivisorsV1(5))
}

func TestFindTrianguleNumberWithNDivisorsV1N500(t *testing.T) {
	assert.Equal(t, 76576500, p012.FindTrianguleNumberWithNDivisorsV1(500))
}

func BenchmarkFindTrianguleNumberWithNDivisorsV1N500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p012.FindTrianguleNumberWithNDivisorsV1(500)
	}
}

func TestFindTrianguleNumberWithNDivisorsV2N5(t *testing.T) {
	assert.Equal(t, 28, p012.FindTrianguleNumberWithNDivisorsV2(5))
}

func TestFindTrianguleNumberWithNDivisorsV2N500(t *testing.T) {
	assert.Equal(t, 76576500, p012.FindTrianguleNumberWithNDivisorsV2(500))
}

func BenchmarkFindTrianguleNumberWithNDivisorsV2N500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p012.FindTrianguleNumberWithNDivisorsV2(500)
	}
}
