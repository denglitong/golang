package leetcode

func trap(height []int) int {
	n := len(height)
	maxI, maxV := 0, height[0]
	leftMax := make([]int, n)
	for i := 0; i < len(height); i++ {
		if height[i] > maxV {
			maxV = height[i]
			maxI = i
		}
		leftMax[i] = maxV
	}
	rightMax := make([]int, n)
	rightMaxV := height[len(height)-1]
	for i := len(height) - 1; i >= maxI; i-- {
		if height[i] > rightMaxV {
			rightMaxV = height[i]
		}
		rightMax[len(height)-1-i] = rightMaxV
	}

	res := 0
	for i := 0; i <= maxI; i++ {
		if leftMax[i] > height[i] {
			res += leftMax[i] - height[i]
		}
	}
	for i := len(height) - 1; i >= maxI; i-- {
		if rightMax[len(height)-1-i] > height[i] {
			res += rightMax[i] - height[i]
		}
	}
	return res
}
