package leetcode

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	mi, mj := make([]int, 0), make([]int, 0)
	len1, len2 := len(matrix), len(matrix[0])
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				mi = append(mi, i)
				mj = append(mj, j)
			}
		}
	}
	for _, i := range mi {
		for j := 0; j < len2; j++ {
			matrix[i][j] = 0
		}
	}
	for _, j := range mj {
		for i := 0; i < len1; i++ {
			matrix[i][j] = 0
		}
	}
}
