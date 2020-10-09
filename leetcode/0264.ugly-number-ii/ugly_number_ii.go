package leetcode

var (
	dp         = append(make([]int, 0, 1690), 0, 1)
	i2, i3, i5 = 1, 1, 1
)

func nthUglyNumber(n int) int {
	if len(dp) > n {
		return dp[n]
	}
	for i := len(dp) - 1; i <= n; i++ {
		uglyNum := min(min(dp[i2]*2, dp[i3]*3), dp[i5]*5)
		dp = append(dp, uglyNum)
		if uglyNum == dp[i2]*2 {
			i2++
		}
		if uglyNum == dp[i3]*3 {
			i3++
		}
		if uglyNum == dp[i5]*5 {
			i5++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
