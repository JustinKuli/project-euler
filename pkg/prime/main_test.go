package prime

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFactorsOf(t *testing.T) {
	tests := []struct {
		in  uint64
		out []uint64
	}{
		{6, []uint64{2, 3}},
		{7, []uint64{7}},
		{10, []uint64{2, 5}},
		{360, []uint64{2, 3, 5}},
		{366, []uint64{2, 3, 61}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Prime Factors of %v", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans, err := FactorsOf(tt.in)
			if err != nil {
				t.Errorf("Unexpected error calculating factors of %v: %v", tt.in, err)
			}
			if !reflect.DeepEqual(ans, tt.out) {
				t.Errorf("Got %v, want %v", ans, tt.out)
			}
		})
	}
}

func TestSliceTo(t *testing.T) {
	tests := []struct {
		in  []uint64
		max uint64
		out []uint64
	}{
		{[]uint64{10, 20, 30, 40}, 25, []uint64{10, 20}},
		{[]uint64{10, 20, 30, 40}, 55, []uint64{10, 20, 30, 40}},
		{[]uint64{10, 20, 30, 40}, 5, []uint64{}},
		{[]uint64{10, 20, 30, 40}, 30, []uint64{10, 20}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v slice to %v", tt.in, tt.max)
		t.Run(testname, func(t *testing.T) {
			ans := sliceTo(tt.in, tt.max)
			if !reflect.DeepEqual(ans, tt.out) {
				t.Errorf("Got %v, want %v", ans, tt.out)
			}
		})
	}
}

func TestCoprime(t *testing.T) {
	tests := []struct {
		in    []uint64
		check uint64
		out   bool
	}{
		{[]uint64{2, 3, 5, 7}, 11, true},
		{[]uint64{2, 3}, 4, false},
		{[]uint64{2, 10}, 20, false},
		{[]uint64{2, 3, 5}, 77, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v coprime to %v", tt.check, tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := coprime(tt.in, tt.check)
			if ans != tt.out {
				t.Errorf("Got %v, want %v", ans, tt.out)
			}
		})
	}
}

func TestIntsqrt(t *testing.T) {
	tests := []struct {
		in  uint64
		out uint64
	}{
		{99, 9},
		{100, 10},
		{101, 10},
		{999_999, 999},
		{1_000_000, 1_000},
		{1_000_001, 1_000},
		{17_748_900_625, 133_225},
		{17_748_900_620, 133_224},
		{17_748_900_630, 133_225},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("integer sqrt of %v", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans, err := intsqrt(tt.in, 1000)
			if err != nil {
				t.Errorf("Got error calculating sqrt of %v: %v", tt.in, err)
			}
			if ans != tt.out {
				t.Errorf("Got %v, want %v", ans, tt.out)
			}
		})
	}
}

func TestAbsDiff(t *testing.T) {
	tests := []struct {
		n1, n2, out uint64
	}{
		{10, 20, 10},
		{20, 10, 10},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("distance between %v and %v", tt.n1, tt.n2)
		t.Run(testname, func(t *testing.T) {
			ans := absDiff(tt.n1, tt.n2)
			if ans != tt.out {
				t.Errorf("Got %v, want %v", ans, tt.out)
			}
		})
	}
}
