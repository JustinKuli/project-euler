package bignum

import (
	"strconv"
)

// Int is a big integer
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

	for j := len(i.n) - 1; j >= 0; j-- {
		digit := strconv.Itoa(int(i.n[j]))
		for len(digit) < bignumDigitCount {
			digit = "0" + digit
		}
		str += digit
	}

	for str[0] == '0' && len(str) > 1 {
		str = str[1:]
	}

	if i.neg {
		str = "-" + str
	}
	return str
}

// Add adds two bignum.Ints
func (i Int) Add(i2 Int) Int {
	// Get the numbers to have the same number of "digits"
	//  by padding with zeroes (on the big side)
	for len(i.n) < len(i2.n) {
		i.n = append(i.n, int32(0))
	}
	for len(i2.n) < len(i.n) {
		i2.n = append(i2.n, int32(0))
	}

	sum := make([]int32, len(i.n))
	kerry := int32(0)
	for id := range i.n {
		sum[id] = i.n[id] + i2.n[id] + kerry
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

// Multiply multiplies two bignum.Ints
func (i Int) Multiply(i2 Int) Int {
	// Ensure that i has fewer digits than i2
	if len(i.n) > len(i2.n) {
		return i2.Multiply(i)
	}

	acc := MakeInt("0")
	for ind := range i.n {
		// Do single "digit" multiplication
		num := i2.copyForMultiply(ind)
		for k, v := range num {
			num[k] = v * int64(i.n[ind])
		}

		// Carry overflows inside single "digit" product
		kerry := int64(0)
		kerrycount := 0
		for idx := range num {
			kerrycount++
			num[idx] += kerry
			if num[idx] > int64(biggestNum) {
				kerry = num[idx] / (int64(biggestNum) + 1)
				num[idx] = num[idx] % (int64(biggestNum) + 1)
			} else {
				kerry = int64(0)
			}
		}

		if kerry != 0 {
			foo := make([]int32, kerrycount)
			foo = append(foo, int32(kerry))
			acc = acc.Add(Int{foo, false})
		}

		// Accumulate
		acc = acc.Add(convertMultiplyBack(num))
	}

	return acc
}

func (i Int) copyForMultiply(pad int) []int64 {
	ans := make([]int64, len(i.n)+pad)
	for k, v := range i.n {
		ans[k+pad] = int64(v)
	}
	return ans
}

func convertMultiplyBack(in []int64) Int {
	ans := make([]int32, len(in))
	for k, v := range in {
		ans[k] = int32(v)
	}
	return Int{ans, false}
}
