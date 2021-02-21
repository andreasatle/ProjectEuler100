package prime_test

import (
	"errors"
	"testing"

	"github.com/andreasatle/ProjectEuler100/prime"
	"github.com/stretchr/testify/assert"
)

func TestComputePrimes20(t *testing.T) {
	primes := prime.ComputePrimes(20)
	assert.Equal(t, 8, len(primes), "Wrong length of slice from prime.ComputePrimes().")
	assert.Equal(t, 2, primes[0], "First prime is wrong.")
	assert.Equal(t, 3, primes[1], "Second prime is wrong.")
	assert.Equal(t, 5, primes[2], "Third prime is wrong.")
	assert.Equal(t, 7, primes[3], "Fourth prime is wrong.")
	assert.Equal(t, 11, primes[4], "Fifth prime is wrong.")
	assert.Equal(t, 13, primes[5], "Sixth prime is wrong.")
	assert.Equal(t, 17, primes[6], "Seventh prime is wrong.")
	assert.Equal(t, 19, primes[7], "Eighth prime is wrong.")
}

func TestComputePrimesNeg(t *testing.T) {
	primes := prime.ComputePrimes(-1)
	assert.Equal(t, 0, len(primes), "Wrong length of slice from prime.ComputePrimes().")
}

func TestComputePrimes1(t *testing.T) {
	primes := prime.ComputePrimes(1)
	assert.Equal(t, 0, len(primes), "Wrong length of slice from prime.ComputePrimes().")
}

func TestComputePrimes2(t *testing.T) {
	primes := prime.ComputePrimes(2)
	assert.Equal(t, 1, len(primes), "Wrong length of slice from prime.ComputePrimes().")
	assert.Equal(t, 2, primes[0], "First prime is wrong.")
}

func TestGeneratePrimes2(t *testing.T) {
	nextPrime := prime.GeneratePrimes(2)
	p, err := nextPrime()
	assert.Equal(t, 2, p, "First prime is wrong.")
	assert.Equal(t, nil, err, "First error message is wrong.")
	p, err = nextPrime()
	assert.Equal(t, errors.New("maxPrime exceeded in prime.GeneratePrimes"), err, "Second error message is wrong.")
}

func TestGeneratePrimes20(t *testing.T) {
	next := prime.GeneratePrimes(20)
	p, err := next()
	assert.Equal(t, 2, p, "First prime is wrong.")
	assert.Equal(t, nil, err, "First error message is wrong.")
	p, err = next()
	assert.Equal(t, 3, p, "Second prime is wrong.")
	assert.Equal(t, nil, err, "Second error message is wrong.")
	p, err = next()
	assert.Equal(t, 5, p, "Third prime is wrong.")
	assert.Equal(t, nil, err, "Third error message is wrong.")
	p, err = next()
	assert.Equal(t, 7, p, "Fourth prime is wrong.")
	assert.Equal(t, nil, err, "Fourth error message is wrong.")
	p, err = next()
	assert.Equal(t, 11, p, "Fifth prime is wrong.")
	assert.Equal(t, nil, err, "Fifth error message is wrong.")
	p, err = next()
	assert.Equal(t, 13, p, "Sixth prime is wrong.")
	assert.Equal(t, nil, err, "Sixth error message is wrong.")
	p, err = next()
	assert.Equal(t, 17, p, "Seventh prime is wrong.")
	assert.Equal(t, nil, err, "Seventh error message is wrong.")
	p, err = next()
	assert.Equal(t, 19, p, "Eighth prime is wrong.")
	assert.Equal(t, nil, err, "Eighth error message is wrong.")
	p, err = next()
	assert.Equal(t, errors.New("maxPrime exceeded in prime.GeneratePrimes"), err, "Second error message is wrong.")
}
