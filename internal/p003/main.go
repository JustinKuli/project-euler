package p003

import (
	"strconv"

	"github.com/JustinKuli/project-euler/pkg/prime"
)

// Solve this problem:
// What is the largest prime factor of the number 600851475143?
func Solve() string {
	factors, err := prime.FactorsOf(uint64(600_851_475_143))
	if err != nil {
		return "Error"
	}
	return strconv.Itoa(int(factors[len(factors)-1]))
}
