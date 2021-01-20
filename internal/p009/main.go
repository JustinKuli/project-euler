package p009

import "strconv"

// Solve this problem:
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
func Solve() string {
	for a := 1; a <= 350; a++ {
		a2 := a * a
		for b := a + 1; b <= 1000; b++ {
			c := 1000 - a - b
			if a2+b*b == c*c {
				return strconv.Itoa(a * b * c)
			}
		}
	}
	return "Error"
}
