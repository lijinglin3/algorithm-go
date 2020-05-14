package leetcode

func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}

	max := 0
	for i := 1; i < l; i++ {
		t := prices[i] - prices[i-1]
		if t > 0 {
			max += t
		}
	}

	return max
}
