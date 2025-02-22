package leetcode

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	i, j := 0, 0
	for {
		for ; i < len(nums) && nums[i] != 0; i++ {
		}
		for j = i + 1; j < len(nums) && nums[j] == 0; j++ {
		}
		if j < len(nums) {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			break
		}
		i++
		j++
	}
}

func moveZeroesV2(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			// if left == right, then this swap change nothing!
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
