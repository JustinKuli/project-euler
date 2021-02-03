package p015

import "strconv"

// Solve this problem:
// Starting in the top left corner of a 2×2 grid, and only being able to move to
// the right and down, there are exactly 6 routes to the bottom right corner.
// How many such routes are there through a 20×20 grid?
func Solve() string {
	return strconv.Itoa(boxPaths(20, 20))
}

type box struct {
	x, y int
}

var knownBoxPaths = make(map[box]int)

func boxPaths(x, y int) int {
	if x == 0 {
		return 1
	}
	if y == 0 {
		return 1
	}
	if x == 1 && y == 1 {
		return 2
	}
	currentBox := box{x, y}
	if ans, ok := knownBoxPaths[currentBox]; ok {
		return ans
	}

	ans := boxPaths(x, y-1) + boxPaths(x-1, y)
	knownBoxPaths[currentBox] = ans
	return ans
}
