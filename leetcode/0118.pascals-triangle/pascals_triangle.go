package leetcode

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	result := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		tmp := []int{1}
		for j := 1; j < i; j++ {
			tmp = append(tmp, result[i-1][j]+result[i-1][j-1])
		}
		tmp = append(tmp, 1)
		result = append(result, tmp)
	}

	return result
}
