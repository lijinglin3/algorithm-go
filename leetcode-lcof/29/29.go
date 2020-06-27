package leetcode

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	lengthStart, widthStart := 0, 0
	lengthEnd, widthEnd := len(matrix), len(matrix[0])
	result := make([]int, 0, lengthEnd*widthEnd)

	count := 0
	for lengthStart < lengthEnd && widthStart < widthEnd {
		switch count % 4 {
		case 0:
			for i := widthStart; i < widthEnd; i++ {
				result = append(result, matrix[lengthStart][i])
			}
			lengthStart++
		case 1:
			for i := lengthStart; i < lengthEnd; i++ {
				result = append(result, matrix[i][widthEnd-1])
			}
			widthEnd--
		case 2:
			for i := widthEnd - 1; i >= widthStart; i-- {
				result = append(result, matrix[lengthEnd-1][i])
			}
			lengthEnd--
		case 3:
			for i := lengthEnd - 1; i >= lengthStart; i-- {
				result = append(result, matrix[i][widthStart])
			}
			widthStart++
		}
		count++
	}

	return result
}
