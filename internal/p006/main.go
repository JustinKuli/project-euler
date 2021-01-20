package p006

import (
	"fmt"
)

// Solve this problem:
// Find the difference between the sum of the squares of the first one hundred
// natural numbers and the square of the sum.
func Solve() {
	sumOfSquares := 0
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		sumOfSquares += i * i
	}
	ans := sum*sum - sumOfSquares
	fmt.Println(ans)
}
