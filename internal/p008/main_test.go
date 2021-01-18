package p008

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
			ans, err := digitalProduct(tt.in)
			if err != nil {
				t.Errorf("Got unexpected error %v", err)
			}
			if ans != tt.out {
				t.Errorf("Got %v, want %v", ans, tt.out)
			}
		})
	}
}
