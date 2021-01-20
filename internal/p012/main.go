package p012

import (
	"strconv"

	"github.com/JustinKuli/project-euler/pkg/prime"
)

// Solve this problem:
// What is the value of the first triangle number to have over five hundred divisors?
func Solve() string {
	triNum := 0
	for i := 1; i < 1_000_000; i++ {
		triNum += i
		if numOfDivisors(triNum) > 500 {
			return strconv.Itoa(triNum)
		}
	}
	return "Error"
}

func numOfDivisors(num int) int {
	pFactors, err := prime.FactorsOf(uint64(num))
	if err != nil {
		panic(err)
	}

	divisors := make(map[uint64]bool)
	for i := 0; i <= (1 << len(pFactors)); i++ {
		divisor := 1
		for j, factor := range pFactors {
			if i&(1<<j) != 0 {
				divisor *= int(factor)
			}
		}
		divisors[uint64(divisor)] = true
	}
	return len(divisors)
}
