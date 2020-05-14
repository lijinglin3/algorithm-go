package leetcode

func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}

	max := 0
	for i := 0; i < l; i++ {
		t := 0
		for j := i; j < l; j++ {
			tmp := prices[j] - prices[i]
			if tmp > t {
				t = tmp
			}
		}
		if t > max {
			max = t
		}
	}

	return max
}
