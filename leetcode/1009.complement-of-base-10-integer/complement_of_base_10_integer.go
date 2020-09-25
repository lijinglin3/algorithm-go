package leetcode

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	i := 1
	for i <= N {
		i <<= 1
	}
	return i - N - 1
}
