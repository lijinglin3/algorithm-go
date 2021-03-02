package leetcode

var cache = []int{0, 1, 1}

func tribonacci(n int) int {
	if n < len(cache) {
		return cache[n]
	}
	ret := tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	cache = append(cache, ret)
	return ret
}
