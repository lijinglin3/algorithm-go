package leetcode

func merge(A []int, m int, B []int, n int) {
	length := m + n - 1
	for m > 0 && n > 0 {
		if A[m-1] > B[n-1] {
			A[length] = A[m-1]
			m--
		} else {
			A[length] = B[n-1]
			n--
		}
		length--
	}
	for ; n > 0; n-- {
		A[n-1] = B[n-1]
	}
}
