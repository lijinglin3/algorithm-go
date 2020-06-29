package leetcode

func kthFactor(n int, k int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
			if count == k {
				return i
			}
		}
	}
	return -1
}
