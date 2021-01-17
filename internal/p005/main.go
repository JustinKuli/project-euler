package p005

import "fmt"

// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?
func Solve() {
	nums := make([]int, 20)
	for i := 1; i <= 20; i++ {
		nums[i-1] = i
	}
	fmt.Println(smallestMultiple(nums))
}

// Returns the smallest multiple of all the numbers in the list.
// Requires the list to be in order from largest to smallest.
func smallestMultiple(nums []int) int {
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
