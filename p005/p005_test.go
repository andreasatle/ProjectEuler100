package p005_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p005"
	"github.com/stretchr/testify/assert"
)

func TestDivisibleUpTo20(t *testing.T) {
	assert.Equal(t, 232792560, p005.DivisibleUpTo20())
}

func BenchmarkDivisibleUpTo20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p005.DivisibleUpTo20()
	}
}
