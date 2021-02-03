package num

import (
	"sort"
	"strconv"
	"strings"

	"github.com/JustinKuli/project-euler/pkg/prime"
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

// DigitalSum returns the sum of all the digits of the input
func DigitalSum(in string) (int, error) {
	sum := 0
	for i := 0; i < len(in); i++ {
		digit, err := strconv.Atoi(string(in[i : i+1]))
		if err != nil {
			return 0, err
		}
		sum += digit
	}
	return sum, nil
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

// Divisors returns the divisors of the input.
func Divisors(num int) []int {
	pFactors, err := prime.FactorsOf(uint64(num))
	if err != nil {
		panic(err)
	}

	// Every divisor is some product of the prime factors.
	// We need to put them in a map, since there could be duplicates.
	// For example, take 60 : prime factors =  2  2  3  5
	//            divisor 6 is the product of  2     3
	//                but it could also be        2  3
	//
	// We can enumerate the divisors (with duplicates) with a binary number.
	// For instance, 6 is either 0b1010 or 0b0110 in the example above.
	divMap := make(map[uint64]bool)
	// reminder: (1 << x) == (2**x)
	for i := 0; i <= (1 << len(pFactors)); i++ {
		divisor := 1
		for j, factor := range pFactors {
			if i&(1<<j) != 0 {
				divisor *= int(factor)
			}
		}
		divMap[uint64(divisor)] = true
	}

	divisors := make([]int, 0, len(divMap))
	for div := range divMap {
		divisors = append(divisors, int(div))
	}

	sort.Ints(divisors)
	return divisors
}
