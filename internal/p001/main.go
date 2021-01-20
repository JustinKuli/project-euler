package p001

import (
	"fmt"
)

// Solve this problem:
// Find the sum of all the multiples of 3 or 5 below 1000.
func Solve() {
	ans := 0
	for i := 1; i < 1000; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			ans += i
		}
	}
	fmt.Println(ans)
}
