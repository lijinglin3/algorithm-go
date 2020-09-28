package leetcode

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch {
			case i == 0 && j == 0:
				continue
			case i == 0:
				grid[i][j] += grid[i][j-1]
			case j == 0:
				grid[i][j] += grid[i-1][j]
			default:
				tmp := grid[i][j-1]
				if grid[i][j-1] < grid[i-1][j] {
					tmp = grid[i-1][j]
				}
				grid[i][j] += tmp
			}
		}
	}
	return grid[m-1][n-1]
}
