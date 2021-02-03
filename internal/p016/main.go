package p016

import (
	"fmt"
	"strconv"

	"github.com/JustinKuli/project-euler/pkg/bignum"
	"github.com/JustinKuli/project-euler/pkg/num"
)

// Solve this problem:
// 2**15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
// What is the sum of the digits of the number 2**1000?
func Solve() string {
	two := bignum.MakeInt("2")
	big := bignum.MakeInt("2")
	for i := 1; i < 1000; i++ {
		big = big.Multiply(two)
	}

	fmt.Println(big)
	ans, _ := num.DigitalSum(big.String())
	return strconv.Itoa(ans)
}
