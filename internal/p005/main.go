package p005

import (
	"strconv"

	"github.com/JustinKuli/project-euler/pkg/num"
)

// Solve this problem:
// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?
func Solve() string {
	nums := make([]int, 20)
	for i := 1; i <= 20; i++ {
		nums[i-1] = i
	}
	return strconv.Itoa(num.SmallestMultiple(nums))
}
