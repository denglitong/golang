package leetcode

// https://leetcode.cn/problems/maximum-subarray
func maxSubArray(nums []int) int {
	maxSums := make([]int, len(nums))
	for _, num := range nums {
		maxSums = append(maxSums, num)
	}

	res := maxSums[0]
	for i := 1; i < len(nums); i++ {
		if maxSums[i-1]+nums[i] > nums[i] {
			maxSums[i] = maxSums[i-1] + nums[i]
		} else {
			maxSums[i] = nums[i]
		}
		if maxSums[i] > res {
			res = maxSums[i]
		}
	}
	return res
}
