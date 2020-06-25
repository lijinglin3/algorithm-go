package leetcode

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	i, j := 0, len(matrix[0])-1
	for {
		switch {
		case matrix[i][j] == target:
			return true
		case matrix[i][j] > target:
			j--
		case matrix[i][j] < target:
			i++
		}
		if i >= len(matrix) || j < 0 {
			break
		}
	}
	return false
}
