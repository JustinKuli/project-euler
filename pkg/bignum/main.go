package bignum

import (
	"strconv"
)

type Int struct {
	n   []int64
	neg bool
}

const bignumDigitCount = 15
const biggestNum = int64(999_999_999_999_999)

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

	numAcc := make([]int64, 0, 0)
	for len(s) > 18 {
		digits, err := strconv.Atoi(s[len(s)-bignumDigitCount:])
		if err != nil {
			panic(err)
		}
		numAcc = append(numAcc, int64(digits))
		s = s[:len(s)-bignumDigitCount]
	}
	digits, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	numAcc = append(numAcc, int64(digits))

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

func (i1 Int) Add(i2 Int) Int {
	// Get the numbers to have the same number of "digits"
	//  by padding with zeroes (on the big side)
	for len(i1.n) < len(i2.n) {
		i1.n = append(i1.n, int64(0))
	}
	for len(i2.n) < len(i1.n) {
		i2.n = append(i2.n, int64(0))
	}

	sum := make([]int64, len(i1.n))
	kerry := int64(0)
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
		sum = append(sum, int64(1))
	}

	return Int{sum, false} // TODO: Handle negatives
}
