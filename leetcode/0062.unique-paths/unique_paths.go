package leetcode

import (
	"math/big"
)

func uniquePaths(m int, n int) int {
	//dp := make([][]int, m)
	//for i := range dp {
	//	dp[i] = make([]int, n)
	//}
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		if i == 0 || j == 0 {
	//			dp[i][j] = 1
	//		} else {
	//			dp[i][j] = dp[i-1][j]+dp[i][j-1]
	//		}
	//	}
	//}
	//return dp[m-1][n-1]

	//dp := make([]int, n)
	//tmp := make([]int, n)
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		if i == 0 || j == 0 {
	//			tmp[j] = 1
	//		} else {
	//			tmp[j] = dp[j] + tmp[j-1]
	//		}
	//	}
	//	dp = tmp
	//}
	//return dp[n-1]

	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
