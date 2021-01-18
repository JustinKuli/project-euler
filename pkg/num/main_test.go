package num

import (
	"fmt"
	"testing"
)

func TestDigitalProduct(t *testing.T) {
	var tests = []struct {
		in  string
		out int
	}{
		{"123", 6},
		{"994560918723", 0},
		{"9989", 5832},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Digital Product of %v", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans, err := DigitalProduct(tt.in)
			if err != nil {
				t.Errorf("Got unexpected error %v", err)
			}
			if ans != tt.out {
				t.Errorf("Got %v, want %v", ans, tt.out)
			}
		})
	}
}

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
			ans := SmallestMultiple(tt.nums)
			if ans != tt.out {
				t.Errorf("Got %v, want %v", ans, tt.out)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		num  int
		pali bool
	}{
		{1000, false},
		{1001, true},
		{999, true},
		{909, true},
		{988, false},
		{1234554321, true},
		{123123123, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.num)
		t.Run(testname, func(t *testing.T) {
			ans := IsPalindrome(tt.num)
			if ans != tt.pali {
				t.Errorf("Got %v, want %v", ans, tt.pali)
			}
		})
	}
}
