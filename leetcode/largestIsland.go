package leetcode

// https://leetcode.cn/problems/making-a-large-island/description/
func largestIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	islandIdx := 2
	island2Area := []int{0, 0} // island area index from 2, put index [0,1] as placeholder
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := searchAndMarkIsland(grid, i, j, islandIdx)
				island2Area = append(island2Area, area)
				islandIdx++
			}
		}
	}

	maxArea := 0
	for i := 2; i < len(island2Area); i++ {
		if island2Area[i] > maxArea {
			maxArea = island2Area[i]
		}
	}
	if maxArea == 0 {
		return maxArea + 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				islands := make(map[int]int)
				// left, up
				leftIslandIdx, leftArea := 0, 0
				if isInGridV2(grid, i, j-1) {
					leftIslandIdx = grid[i][j-1]
					leftArea = island2Area[leftIslandIdx]
					maxArea = max(maxArea, leftArea+1)
				}
				islands[leftIslandIdx] = leftArea

				upIslandIdx, upArea := 0, 0
				if isInGridV2(grid, i-1, j) {
					upIslandIdx = grid[i-1][j]
					upArea = island2Area[upIslandIdx]
					maxArea = max(maxArea, upArea+1)
				}
				islands[upIslandIdx] = upArea

				rightIslandIdx, rightArea := 0, 0
				if isInGridV2(grid, i, j+1) {
					rightIslandIdx = grid[i][j+1]
					rightArea = island2Area[rightIslandIdx]
					maxArea = max(maxArea, rightArea+1)
				}
				islands[rightIslandIdx] = rightArea

				downIslandIdx, downArea := 0, 0
				if isInGridV2(grid, i+1, j) {
					downIslandIdx = grid[i+1][j]
					downArea = island2Area[downIslandIdx]
					maxArea = max(maxArea, downArea+1)
				}
				islands[downIslandIdx] = downArea

				area := 1
				for _, v := range islands {
					area += v
				}
				maxArea = max(maxArea, area)
			}
		}
	}

	//fmt.Println(grid)
	//fmt.Println("islandIdxToArea:", island2Area)
	//fmt.Println(maxArea)
	return maxArea
}

func searchAndMarkIsland(grid [][]int, i, j int, islandIdx int) int {
	if !isInGridV2(grid, i, j) {
		return 0
	}
	if grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = islandIdx
	return 1 +
		searchAndMarkIsland(grid, i-1, j, islandIdx) +
		searchAndMarkIsland(grid, i+1, j, islandIdx) +
		searchAndMarkIsland(grid, i, j-1, islandIdx) +
		searchAndMarkIsland(grid, i, j+1, islandIdx)
}

func isInGridV2(grid [][]int, i, j int) bool {
	return (0 <= i && i < len(grid)) && (0 <= j && j < len(grid[0]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
