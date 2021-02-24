package p013_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p013"
	"github.com/stretchr/testify/assert"
)

func TestSumNumberSlice(t *testing.T) {
	sum := p013.SumNumberSlice(p013.Input)
	assert.Equal(t, "5537376230", sum[:10])
}
