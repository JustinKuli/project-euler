package p014

import "strconv"

// Solve this problem:
// The following iterative sequence is defined for the set of positive integers:
// 		n → n/2 (n is even)
// 		n → 3n + 1 (n is odd)
// Which starting number, under one million, produces the longest chain?
func Solve() string {
	bigStartNum := 0
	bigLength := 0
	for i := 1; i < 1_000_000; i++ {
		length := collatzLength(i)
		if length > bigLength {
			bigStartNum = i
			bigLength = length
		}
	}
	return strconv.Itoa(bigStartNum)
}

var calced map[int]int = make(map[int]int)

func collatzLength(num int) int {
	if num == 1 {
		return 1
	}

	if ans, ok := calced[num]; ok {
		return ans
	}

	if num%2 == 0 {
		calced[num] = collatzLength(num/2) + 1
	} else {
		calced[num] = collatzLength(num*3+1) + 1
	}
	return calced[num]
}
