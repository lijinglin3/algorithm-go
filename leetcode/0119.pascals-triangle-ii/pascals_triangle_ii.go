package leetcode

func getRow(rowIndex int) []int {
	result := []int{1}
	if rowIndex <= 0 {
		return result
	}
	for i := 1; i <= rowIndex; i++ {
		tmp := make([]int, 0, i+1)
		tmp = append(tmp, 1)
		for j := 1; j < i; j++ {
			tmp = append(tmp, result[j-1]+result[j])
		}
		tmp = append(tmp, 1)
		result = tmp
	}
	return result
}
