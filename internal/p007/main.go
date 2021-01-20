package p007

import (
	"fmt"

	"github.com/JustinKuli/project-euler/pkg/prime"
)

// Solve this problem:
// What is the 10 001st prime number?
func Solve() {
	size := uint64(1_000_000)
	primes := prime.List(size)
	for len(primes) < 10_001 {
		size *= 2
		primes = prime.List(size)
	}
	fmt.Println(primes[10_000])
}
