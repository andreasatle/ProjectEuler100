package p008_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p008"
	"github.com/stretchr/testify/assert"
)

func TestMaxFactor4(t *testing.T) {
	assert.Equal(t, 5832, p008.MaxFactor(4))
}

func TestMaxFactor13(t *testing.T) {
	assert.Equal(t, 23514624000, p008.MaxFactor(13))
}

func BenchmarkMaxFactor13(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p008.MaxFactor(13)
	}
}
