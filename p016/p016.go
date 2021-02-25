package p016

import (
	"math/big"
)

// PowerDigitSum returns the digit sum of base^pow.
func PowerDigitSum(base, pow int) int {
	zeroBig := &big.Int{}
	baseBig := big.NewInt(int64(base))
	powBig := big.NewInt(int64(pow))
	splitBig := big.NewInt(int64(1000000000000000000))
	zBig := &big.Int{}
	mBig := &big.Int{}
	zBig = zBig.Exp(baseBig, powBig, mBig)

	digitSum := 0
	for zBig.Cmp(zeroBig) == 1 {
		zBig.DivMod(zBig, splitBig, mBig)
		digitSum += computeDigitSum(mBig.Int64())
	}

	return digitSum
}

func computeDigitSum(n int64) int {
	nn := int(n)
	sum := 0
	for ; nn > 0; nn /= 10 {
		sum += nn % 10
	}
	return sum
}
