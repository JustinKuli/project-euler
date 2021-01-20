package p007

import (
	"strconv"

	"github.com/JustinKuli/project-euler/pkg/prime"
)

// Solve this problem:
// What is the 10 001st prime number?
func Solve() string {
	size := uint64(1_000_000)
	primes := prime.List(size)
	for len(primes) < 10_001 {
		size *= 2
		primes = prime.List(size)
	}
	return strconv.Itoa(int(primes[10_000]))
}
