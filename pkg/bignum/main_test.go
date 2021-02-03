package bignum

import (
	"fmt"
	"testing"
)

func TestIntAdd(t *testing.T) {
	var tests = []struct {
		num1 Int
		num2 Int
		sum  Int
	}{
		{MakeInt("1000"), MakeInt("1000"), MakeInt("2000")},
		{MakeInt("0"), MakeInt("100000000000000"), MakeInt("100000000000000")},
		{
			MakeInt("111111111111111111111"),
			MakeInt("111111111111111111111"),
			MakeInt("222222222222222222222"),
		},
		{
			MakeInt("444444"),
			MakeInt("444444"),
			MakeInt("888888"),
		},
		{
			MakeInt("222222222222"),
			MakeInt("222222222222"),
			MakeInt("444444444444"),
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v + %v", tt.num1, tt.num2)
		t.Run(testname, func(t *testing.T) {
			ans := tt.num1.Add(tt.num2)
			if ans.String() != tt.sum.String() {
				t.Errorf("Got %v, want %v", ans, tt.sum)
			}
		})
	}
}

func TestIntString(t *testing.T) {
	var tests = []struct {
		num Int
		str string
	}{
		{MakeInt("0"), "0"},
		{MakeInt("1000"), "1000"},
		{MakeInt("1000000000000"), "1000000000000"},
		{MakeInt("12345678901"), "12345678901"},
		{MakeInt("123456789012345678901234567890"), "123456789012345678901234567890"},
		{MakeInt("-123456789012345678901234567890"), "-123456789012345678901234567890"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.num)
		t.Run(testname, func(t *testing.T) {
			ans := tt.num.String()
			if ans != tt.str {
				t.Errorf("Got %v, want %v", ans, tt.str)
			}
		})
	}
}

func TestIntMultiply(t *testing.T) {
	var tests = []struct {
		i, j Int
		str  string
	}{
		{MakeInt("100"), MakeInt("100"), "10000"},
		{MakeInt("5000000"), MakeInt("5000000"), "25000000000000"},
		{MakeInt("2000000000"), MakeInt("2"), "4000000000"},
		{MakeInt("2"), MakeInt("2000000000"), "4000000000"},
		{MakeInt("3000000000"), MakeInt("3000000000"), "9000000000000000000"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v * %v", tt.i, tt.j)
		t.Run(testname, func(t *testing.T) {
			ans := tt.i.Multiply(tt.j).String()
			if ans != tt.str {
				t.Errorf("Got %v, want %v", ans, tt.str)
			}
		})
	}
}
