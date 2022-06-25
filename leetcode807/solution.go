package leetcode807

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rowMax := make([]int, n)
	colMax := make([]int, n)
	for i := 0; i < n; i++ {
		row := 0
		col := 0
		for j := 0; j < n; j++ {
			if grid[i][j] > row {
				row = grid[i][j]
			}
			if grid[j][i] > col {
				col = grid[j][i]
			}
		}
		rowMax[i] = row
		colMax[i] = col
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += min(rowMax[i], colMax[j]) - grid[i][j]
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
