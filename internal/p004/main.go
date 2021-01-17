package p004

import (
	"fmt"
	"strconv"
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
			if isPalindrome(prod) {
				largest = prod
			}
		}
	}
	fmt.Println(largest)
}

// Returns whether the input number is a palindrome (in base 10).
func isPalindrome(num int) bool {
	str := strconv.Itoa(num)
	last := len(str) - 1
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[last-i] {
			return false
		}
	}
	return true
}
