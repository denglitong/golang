package leetcode

// https://leetcode.cn/problems/first-missing-positive/
// 用数组进行原地哈希
func firstMissingPositive(nums []int) int {
	n := len(nums)
	nums = append(nums, -1)
	for i, v := range nums {
		if v < 0 || v > n {
			nums[i] = -1 // out of range, number i not occurs
		} else {
			for v >= 0 && v <= n && nums[v] != v {
				nv := nums[v]
				nums[v] = v // within range, number v occurs
				v = nv
			}
		}
	}
	// fmt.Println(nums)
	i := 1
	for ; i < n+1; i++ {
		if nums[i] != i {
			return i
		}
	}
	return i
}
