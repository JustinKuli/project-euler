package num

import (
	"strconv"
	"strings"
)

// DigitalProduct returns the product of all the digits of the input
func DigitalProduct(in string) (int, error) {
	if strings.Contains(in, "0") {
		return 0, nil
	}

	prod := 1
	for i := 0; i < len(in); i++ {
		digit, err := strconv.Atoi(string(in[i : i+1]))
		if err != nil {
			return 0, err
		}
		prod *= digit
	}
	return prod, nil
}

// SmallestMultiple returns the smallest multiple of all the numbers in the list.
// Requires the list to be in order from largest to smallest.
func SmallestMultiple(nums []int) int {
	step := nums[0]
	ans := nums[0]
	for _, i := range nums[1:] {
		for ans%i != 0 {
			ans += step
		}
		step = ans
	}
	return ans
}

// IsPalindrome returns whether the input number is a palindrome (in base 10).
func IsPalindrome(num int) bool {
	str := strconv.Itoa(num)
	last := len(str) - 1
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[last-i] {
			return false
		}
	}
	return true
}
