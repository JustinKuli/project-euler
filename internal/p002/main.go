package p002

import "strconv"

// Solve this problem:
// By considering the terms in the Fibonacci sequence whose values do not exceed
// four million, find the sum of the even-valued terms.
func Solve() string {
	sum := 2
	fib1, fib2 := 1, 2
	for {
		fib1 = fib1 + fib2
		if fib1 > 4_000_000 {
			break
		}
		if fib1%2 == 0 {
			sum += fib1
		}

		fib2 = fib1 + fib2
		if fib2 > 4_000_000 {
			break
		}
		if fib2%2 == 0 {
			sum += fib2
		}
	}
	return strconv.Itoa(sum)
}
