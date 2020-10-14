package leetcode

func numWays(n int, k int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	if n <= 2 {
		return k * k
	}
	slow := k
	fast := k * k
	for i := 3; i <= n; i++ {
		fast, slow = slow*(k-1)+fast*(k-1), fast
	}
	return fast
}
