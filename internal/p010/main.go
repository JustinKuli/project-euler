package p010

import (
	"strconv"

	"github.com/JustinKuli/project-euler/pkg/prime"
)

// Solve this problem:
// Find the sum of all the primes below two million.
func Solve() string {
	primes := prime.List(2_000_000)
	sum := uint64(0)
	for _, p := range primes {
		sum += p
	}
	return strconv.Itoa(int(sum))
}
