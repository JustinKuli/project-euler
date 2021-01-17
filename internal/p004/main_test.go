package p004

import (
	"fmt"
	"testing"
)

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
			ans := isPalindrome(tt.num)
			if ans != tt.pali {
				t.Errorf("Got %v, want %v", ans, tt.pali)
			}
		})
	}
}
