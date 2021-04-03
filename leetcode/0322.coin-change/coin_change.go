package leetcode

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt64
		for _, c := range coins {
			if i-c >= 0 && dp[i-c] != math.MaxInt64 {
				if dp[i] > dp[i-c]+1 {
					dp[i] = dp[i-c] + 1
				}
			}
		}
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}
