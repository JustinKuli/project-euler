package p012

import (
	"strconv"

	"github.com/JustinKuli/project-euler/pkg/num"
)

// Solve this problem:
// What is the value of the first triangle number to have over five hundred divisors?
func Solve() string {
	triNum := 0
	for i := 1; i < 1_000_000; i++ {
		triNum += i
		if num.CountDivisors(triNum) > 500 {
			return strconv.Itoa(triNum)
		}
	}
	return "Error"
}
