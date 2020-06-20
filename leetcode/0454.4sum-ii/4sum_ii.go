package leetcode

func fourSumCount(A []int, B []int, C []int, D []int) int {
	tmp := make(map[int]int)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			tmp[A[i]+B[j]]++
		}
	}

	sum := 0
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			sum += tmp[-C[i]-D[j]]
		}
	}

	return sum
}
