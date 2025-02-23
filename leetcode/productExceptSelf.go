package leetcode

// https://leetcode.cn/problems/product-of-array-except-self/description
func productExceptSelf(nums []int) []int {
	n := len(nums)
	leftProduct, rightProduct := make([]int, n), make([]int, n)
	leftProduct[0], rightProduct[0] = 1, 1
	for i := 1; i < n; i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
		rightProduct[i] = rightProduct[i-1] * nums[n-i]
	}
	// fmt.Println(leftProduct)

	res := []int{}
	for i := 0; i < n; i++ {
		res = append(res, leftProduct[i]*rightProduct[n-1-i])
	}
	return res
}
