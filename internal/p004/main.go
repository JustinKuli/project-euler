package p004

import (
	"fmt"

	"github.com/JustinKuli/project-euler/pkg/num"
)

// Find the largest palindrome made from the product of two 3-digit numbers.
func Solve() {
	largest := 0
	for i := 100; i <= 999; i++ {
		for j := 100; j <= i; j++ {
			prod := i * j
			if prod < largest {
				continue
			}
			if num.IsPalindrome(prod) {
				largest = prod
			}
		}
	}
	fmt.Println(largest)
}
