package leetcode

func numWays(n int) int {
	if n <= 1 {
		return 1
	}

	dp := []int{0, 1, 2}
	for i := 3; i <= n; i++ {
		dp = append(dp, (dp[i-1]+dp[i-2])%1000000007)
	}
	return dp[n]
}
