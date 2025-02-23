package leetcode

// https://leetcode.cn/problems/subarray-sum-equals-k/description/
func subarraySum(nums []int, k int) int {
	n := len(nums)
	leftSums := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			leftSums[i] = nums[i]
		} else {
			leftSums[i] += leftSums[i-1] + nums[i]
		}
	}

	leftSumsMap := make(map[int][]int)
	for i, leftSum := range leftSums {
		if _, ok := leftSumsMap[leftSum]; !ok {
			leftSumsMap[leftSum] = []int{i}
		} else {
			leftSumsMap[leftSum] = append(leftSumsMap[leftSum], i)
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		if leftSums[i] == k {
			res++
		}
		if idx, ok := leftSumsMap[k-leftSums[i]]; ok {
			j, length := 0, len(idx)
			for ; j < length && idx[j] <= i; j++ {
			}
			if j < length {
				res += length - j
			}
		}
	}
	return res
}
