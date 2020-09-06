package leetcode

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	var count int
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			if grid[x][y] == '1' {
				count++
				dfs(x, y, grid)
			}
		}
	}
	return count
}

func dfs(x, y int, grid [][]byte) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	dfs(x-1, y, grid)
	dfs(x+1, y, grid)
	dfs(x, y-1, grid)
	dfs(x, y+1, grid)
}
