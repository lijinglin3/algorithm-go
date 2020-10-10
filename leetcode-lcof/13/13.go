package leetcode

func movingCount(m int, n int, k int) int {
	array := make([][]int, m)
	for i := range array {
		array[i] = make([]int, n)
	}

	array[0][0] = 1
	result := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i > 0 && array[i-1][j] == 1) || (j > 0 && array[i][j-1] == 1) {
				if check(i, j, k) {
					array[i][j] = 1
					result++
				}
			}
		}
	}
	return result
}

func check(m, n, k int) bool {
	res := 0
	for m > 0 {
		res += m % 10
		m /= 10
	}
	for n > 0 {
		res += n % 10
		n /= 10
	}

	return res <= k
}
