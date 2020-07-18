package leetcode

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}
func dfs(board [][]byte, i, j, k int, word string) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return false
	}

	if board[i][j] == word[k] {
		temp := board[i][j]
		board[i][j] = ' '
		if dfs(board, i, j+1, k+1, word) ||
			dfs(board, i, j-1, k+1, word) ||
			dfs(board, i+1, j, k+1, word) ||
			dfs(board, i-1, j, k+1, word) {
			return true
		}
		board[i][j] = temp
	}

	return false
}
