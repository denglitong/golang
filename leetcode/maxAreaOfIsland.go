package leetcode

// https://leetcode.cn/problems/max-area-of-island/
func maxAreaOfIsland(grid [][]int) int {
	res, m, n := 0, len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := gridArea(grid, i, j)
				if area > res {
					res = area
				}
			}
		}
	}
	return res
}

func gridArea(grid [][]int, i, j int) int {
	if !inGrid(grid, i, j) {
		return 0
	}
	if grid[i][j] != 1 {
		return 0
	}
	// mark as visited, avoid visit repeatable
	grid[i][j] = 2
	return 1 +
		gridArea(grid, i-1, j) +
		gridArea(grid, i+1, j) +
		gridArea(grid, i, j-1) +
		gridArea(grid, i, j+1)
}

func inGrid(grid [][]int, i, j int) bool {
	return (0 <= i && i < len(grid)) && (0 <= j && j < len(grid[0]))
}
