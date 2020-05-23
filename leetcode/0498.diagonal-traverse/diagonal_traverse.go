package leetcode

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	length, width := len(matrix), len(matrix[0])
	result := make([]int, 0, length*width)
	up := true
	i, j := 0, 0
	for i < length || j < width {
		if up {
			for i >= 0 && j < width {
				result = append(result, matrix[i][j])
				i--
				j++
			}
			if j <= width-1 {
				i++
			} else {
				i += 2
				j--
			}
			up = false
		} else {
			for i < length && j >= 0 {
				result = append(result, matrix[i][j])
				i++
				j--
			}
			if i <= length-1 {
				j++
			} else {
				j += 2
				i--
			}
			up = true
		}
	}
	return result
}
