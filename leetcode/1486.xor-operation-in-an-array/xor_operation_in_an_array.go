package leetcode

func xorOperation(n int, start int) int {
	res := start
	for i := 1; i < n; i++ {
		res ^= start + 2*i
	}
	return res
}
