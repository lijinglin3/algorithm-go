package leetcode

import "math"

func maxProfit(prices []int) int {
	cost, profit := math.MaxInt64, 0
	for i := 0; i < len(prices); i++ {
		if cost > prices[i] {
			cost = prices[i]
		}
		if delta := prices[i] - cost; profit < delta {
			profit = delta
		}
	}
	return profit
}
