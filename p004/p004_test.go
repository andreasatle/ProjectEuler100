package p004_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p004"
	"github.com/stretchr/testify/assert"
)

func TestLargestPalindrome(t *testing.T) {
	palin, _, _ := p004.LargestPalindrom(100, 999)
	assert.Equal(t, 906609, palin)
}

func BenchmarkLargestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p004.LargestPalindrom(100, 999)
	}
}
