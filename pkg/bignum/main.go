package bignum

import (
	"strconv"
)

type Int struct {
	n   []int32
	neg bool
}

const bignumDigitCount = 9
const biggestNum = int32(999_999_999)

// MaxUint64 = 1<<64 - 1 :: 18_446_744_073_709_551_615
// MaxInt64  = 1<<63 - 1 ::  9_223_372_036_854_775_807
// MaxUint32 = 1<<32 - 1 ::              4_294_967_295
// MaxInt32  = 1<<31 - 1 ::              2_147_483_647

// MakeInt makes a bignum.Int out of the input string.
func MakeInt(s string) Int {
	// Get the first few digits in order to determine the sign of the number
	l := len(s)
	if l > 3 {
		l = 3
	}
	sgnCheck, err := strconv.Atoi(s[:l])
	if err != nil {
		panic(err)
	}
	if sgnCheck < 0 {
		// remove the negative sign
		s = s[1:]
	}

	numAcc := make([]int32, 0, 0)
	for len(s) > bignumDigitCount {
		digits, err := strconv.Atoi(s[len(s)-bignumDigitCount:])
		if err != nil {
			panic(err)
		}
		numAcc = append(numAcc, int32(digits))
		s = s[:len(s)-bignumDigitCount]
	}
	digits, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	numAcc = append(numAcc, int32(digits))

	return Int{
		numAcc,
		sgnCheck < 0,
	}
}

func (i Int) String() string {
	str := ""
	if i.neg {
		str = "-"
	}

	for j := len(i.n) - 1; j >= 0; j-- {
		str += strconv.Itoa(int(i.n[j]))
	}

	// TODO: Pull out 0's?
	return str
}

// Add adds two bignum.Ints
func (i1 Int) Add(i2 Int) Int {
	// Get the numbers to have the same number of "digits"
	//  by padding with zeroes (on the big side)
	for len(i1.n) < len(i2.n) {
		i1.n = append(i1.n, int32(0))
	}
	for len(i2.n) < len(i1.n) {
		i2.n = append(i2.n, int32(0))
	}

	sum := make([]int32, len(i1.n))
	kerry := int32(0)
	for id := range i1.n {
		sum[id] = i1.n[id] + i2.n[id] + kerry
		if sum[id] > biggestNum {
			sum[id] -= (biggestNum + 1)
			kerry = 1
		} else {
			kerry = 0
		}
	}

	if kerry == 1 {
		sum = append(sum, int32(1))
	}

	return Int{sum, false} // TODO: Handle negatives
}
