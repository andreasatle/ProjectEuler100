package p016_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p016"
	"github.com/stretchr/testify/assert"
)

func TestPowerDigitSum(t *testing.T) {
	assert.Equal(t, 26, p016.PowerDigitSum(2, 15))
	assert.Equal(t, 1366, p016.PowerDigitSum(2, 1000))
}

func BenchmarkPowerDigitSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p016.PowerDigitSum(2, 1000)
	}
}
