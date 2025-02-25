package leetcode

// https://leetcode.cn/problems/search-insert-position/description/
func searchInsert(nums []int, target int) int {
	l, r, m := 0, len(nums)-1, 0
	for l < r {
		m = l + (r-l+1)/2
		//fmt.Println(l, r, m)
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m
		}
	}
	l = 0
	for ; l < len(nums) && nums[l] < target; l++ {
	}
	return l
}
