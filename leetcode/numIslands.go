package leetcode

// https://leetcode.cn/problems/number-of-islands/description/
func numIslands(grid [][]byte) int {
	res, m, n := 0, len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				markIsland(grid, i, j)
			}
		}
	}
	return res
}

func markIsland(grid [][]byte, i, j int) {
	if !isInGrid(grid, i, j) {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	markIsland(grid, i-1, j)
	markIsland(grid, i+1, j)
	markIsland(grid, i, j-1)
	markIsland(grid, i, j+1)
}

func isInGrid(grid [][]byte, i, j int) bool {
	return (0 <= i && i < len(grid)) && (0 <= j && j < len(grid[0]))
}
