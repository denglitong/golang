package leetcode

import "testing"

func TestLongestConsecutive(t *testing.T) {
	longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	longestConsecutive([]int{1, 0, 1, 2})
	longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6})
}
