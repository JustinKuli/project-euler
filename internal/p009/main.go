package p009

import (
	"fmt"
)

// Solve this problem:
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
func Solve() {
	for a := 1; a <= 350; a++ {
		a2 := a * a
		for b := a + 1; b <= 1000; b++ {
			c := 1000 - a - b
			if a2+b*b == c*c {
				fmt.Println(a * b * c)
				return
			}
		}
	}
}
