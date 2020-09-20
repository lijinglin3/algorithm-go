package leetcode

var cache = map[int]int{0: 1}

func numTrees(n int) int {
	if v, ok := cache[n]; ok {
		return v
	}
	cache[n] = numTrees(n-1) * 2 * (2*n - 1) / (n + 1)
	return cache[n]
}
