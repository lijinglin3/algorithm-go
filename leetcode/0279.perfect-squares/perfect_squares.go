package leetcode

var (
	sum = make([]int, 0, 100)
	dp  = append(make([]int, 0, 1000), 0)
)

func numSquares(n int) int {
	for i := len(sum); i <= n/2+2; i++ {
		sum = append(sum, i*i)
	}

	for i := len(dp); i <= n; i++ {
		mid, j := n, 1
		for ; i >= sum[j]; j++ {
			mid = min(dp[i-sum[j]]+1, mid)

		}
		dp = append(dp, mid)
	}
	return dp[n]
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
