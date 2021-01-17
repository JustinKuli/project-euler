package p005

import (
	"fmt"
	"testing"
)

func TestSmallestMultiple(t *testing.T) {
	var tests = []struct {
		nums []int
		out  int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2520},
		{[]int{2, 5}, 10},
		{[]int{1, 2, 3, 4, 5, 6, 10, 12, 15, 20, 30, 60}, 60},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("smallest multiple of %v", tt.nums)
		t.Run(testname, func(t *testing.T) {
			ans := smallestMultiple(tt.nums)
			if ans != tt.out {
				t.Errorf("Got %v, want %v", ans, tt.out)
			}
		})
	}
}
