package leetcode

func trailingZeroes(n int) int {
	res := 0
	for n > 0 {
		n /= 5
		res += n
	}
	return res
}
