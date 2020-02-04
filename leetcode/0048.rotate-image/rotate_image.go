package leetcode

func Rotate(matrix [][]int) {
	length := len(matrix)
	if length <= 1 {
		return
	}

	for times := 0; times <= length/2; times++ {
		l := length - 2*times
		for i := 0; i < l-1; i++ {
			tmp := matrix[times][times+i]
			matrix[times][times+i] = matrix[l+times-i-1][times]
			matrix[l+times-i-1][times] = matrix[l+times-1][l+times-i-1]
			matrix[l+times-1][l+times-i-1] = matrix[times+i][l+times-1]
			matrix[times+i][l+times-1] = tmp
		}
	}
}
