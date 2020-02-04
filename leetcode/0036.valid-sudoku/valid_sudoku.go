package leetcode

func IsValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]interface{}, 9)
	columns := make([]map[byte]interface{}, 9)
	boxes := make([]map[byte]interface{}, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]interface{})
		columns[i] = make(map[byte]interface{})
		boxes[i] = make(map[byte]interface{})
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if string(board[i][j]) == "." {
				continue
			}

			t := i/3*3 + j/3
			if _, exist := boxes[t][board[i][j]]; exist {
				return false
			} else {
				boxes[t][board[i][j]] = struct{}{}
			}

			if _, exist := rows[i][board[i][j]]; exist {
				return false
			} else {
				rows[i][board[i][j]] = struct{}{}
			}

			if _, exist := columns[j][board[i][j]]; exist {
				return false
			} else {
				columns[j][board[i][j]] = struct{}{}
			}
		}
	}

	return true
}
