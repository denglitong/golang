package leetcode

import "slices"

// https://leetcode.cn/problems/merge-intervals/description
func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	res := make([][]int, 0)
	for _, interval := range intervals {
		if len(res) == 0 {
			res = append(res, interval)
		} else if interval[0] <= res[len(res)-1][1] {
			if interval[1] > res[len(res)-1][1] {
				res[len(res)-1][1] = interval[1]
			}
		} else {
			res = append(res, interval)
		}
	}
	return res
}
