package p005

import (
	"fmt"

	"github.com/JustinKuli/project-euler/pkg/num"
)

// Solve this problem:
// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?
func Solve() {
	nums := make([]int, 20)
	for i := 1; i <= 20; i++ {
		nums[i-1] = i
	}
	fmt.Println(num.SmallestMultiple(nums))
}
