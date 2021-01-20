package prime

import "fmt"

// FactorsOf returns the prime factors of the given number
func FactorsOf(n uint64) ([]uint64, error) {
	fmt.Println("input:", n)
	sqrt, err := intsqrt(n, 10_000)
	if err != nil {
		return nil, fmt.Errorf("Took too long to calculate sqrt(%v)", n)
	}

	// intsqrt is always less than or equal to the actual sqrt...
	// But we need to be greater than or equal to the actual sqrt
	sqrt++
	fmt.Println("Adjusted sqrt:", sqrt)

	factors := make([]uint64, 0)
	primes := List(sqrt)
	fmt.Println("using primes:", primes)
	for _, p := range primes {
		for n%p == 0 {
			factors = append(factors, p)
			n /= p
		}
		if n == 1 {
			break
		}
	}

	// Handle case where the input is prime
	if n != 1 {
		factors = append(factors, n)
	}
	return factors, nil
}

// TODO: update List to use a sieve for efficiency

// List returns the prime numbers less than or equal to the given `max`
func List(max uint64) []uint64 {
	// Pre-populate the list up to my favorite prime.
	primes := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}
	if max <= primes[len(primes)-1] {
		return sliceTo(primes, max)
	}

	i := primes[len(primes)-1] + 2
	for i <= max {
		if coprime(primes, i) {
			primes = append(primes, i)
		}
		i += 2
	}

	return primes
}

// Returns the sub-slice of values in `input` up to and incuding the given max.
// Requires the input slice to be sorted.
func sliceTo(input []uint64, max uint64) []uint64 {
	for i, val := range input {
		if val > max {
			return input[:i]
		}
	}
	return input
}

// Returns whether the given number is co-prime to all values in the given list.
// Requires the input slice to be sorted.
// Requires all values in the input slice to be prime numbers.
func coprime(list []uint64, check uint64) bool {
	for _, val := range list {
		if val*val > check { // TODO: fix possible overflow
			// We only need to check primes up to the square root
			return true
		}
		if check%val == 0 {
			return false
		}
	}
	return true
}

// Caclulates the integer square root of `n`
// Uses Newton's method to solve `x^2 - n = 0`
func intsqrt(n uint64, maxIters int) (uint64, error) {
	x0 := uint64(0)
	x1 := n
	for i := 0; i <= maxIters/2; i++ {
		x0 = (x1 + n/x1) / 2 // TODO: determine if overflow is possible here
		if absDiff(x0, x1) <= 2 {
			// Stops if the delta is less than 2, to prevent a possible loop
			return x0, nil
		}

		x1 = (x0 + n/x0) / 2 // TODO: determine if overflow is possible here
		if absDiff(x0, x1) <= 2 {
			return x1, nil
		}
	}
	return x1, fmt.Errorf("Exceeded desired iterations (%v)", maxIters)
}

// Returns the absolute value of the difference between the inputs.
func absDiff(n1, n2 uint64) uint64 {
	if n1 > n2 {
		return n1 - n2
	}
	return n2 - n1
}
