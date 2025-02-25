package leetcode

// https://leetcode.cn/problems/permutations/
func permute(nums []int) [][]int {
	/*
	   [0] [1]
	       [0,1]
	       [1,0]
	   [1] [2]
	       [1,2]
	       [2,1]
	           [3]
	           [3,1,2] [1,3,2] [1,2,3]
	           [3,2,1] [2,3,1] [2,1,3]
	*/
	n := len(nums)
	if n == 1 {
		return [][]int{nums}
	}
	prev := permute(nums[:n-1])
	res := [][]int{}
	for _, arr := range prev {
		for i := 0; i < len(arr); i++ {
			item := append([]int{}, arr[0:i]...)
			item = append(item, nums[n-1])
			item = append(item, arr[i:]...)
			// arr[:i] arr[i:]
			res = append(res, item)
		}
		res = append(res, append(arr, nums[n-1]))
	}
	return res
}
