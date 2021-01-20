package p003

import (
	"fmt"

	"github.com/JustinKuli/project-euler/pkg/prime"
)

// Solve this problem:
// What is the largest prime factor of the number 600851475143?
func Solve() {
	factors, err := prime.FactorsOf(uint64(600_851_475_143))
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(factors[len(factors)-1])
}
