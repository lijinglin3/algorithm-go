package leetcode

func finalPrices(prices []int) []int {
	result := make([]int, 0, len(prices))
	flag := false
	for i, v := range prices {
		for _, tmp := range prices[i+1:] {
			if tmp <= v {
				result = append(result, v-tmp)
				flag = true
				break
			}
		}
		if !flag {
			result = append(result, v)
		}
		flag = false
	}
	return result
}
