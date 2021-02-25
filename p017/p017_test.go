package p017_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p017"
	"github.com/stretchr/testify/assert"
)

func TestLetterCountV1(t *testing.T) {
	assert.Equal(t, 21124, p017.LetterCountV1())
}

func TestLetterCountV2(t *testing.T) {
	assert.Equal(t, 21124, p017.LetterCountV2())
}

func BenchmarkLetterCountV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p017.LetterCountV1()
	}
}

func BenchmarkLetterCountV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p017.LetterCountV2()
	}
}
