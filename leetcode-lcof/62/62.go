package leetcode

func lastRemaining(n int, m int) int {
	result := 0
	for i := 2; i <= n; i++ {
		result = (m + result) % i
	}
	return result
}
