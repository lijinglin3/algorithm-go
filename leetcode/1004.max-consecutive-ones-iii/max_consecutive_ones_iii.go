package leetcode

func longestOnes(A []int, K int) int {
	left, right, count := 0, 0, 0
	for ; right < len(A); right++ {
		if A[right] == 0 {
			count++
		}
		if count > K {
			if A[left] == 0 {
				count--
			}
			left++
		}

	}
	return right - left
}
