package leetcode

import "testing"

func TestLargestIsland(t *testing.T) {
	largestIsland([][]int{
		{1, 0},
		{0, 1},
	})
	largestIsland([][]int{
		{1, 1},
		{1, 0},
	})
	largestIsland([][]int{
		{0},
	})
	largestIsland([][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
		{0, 1, 1, 1, 1, 0, 0},
	})
}
