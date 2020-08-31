package leetcode

import "math"

func divide(dividend int, divisor int) int {
	if divisor == -1 && dividend == -(math.MaxInt32+1) {
		return math.MaxInt32
	}
	dividendAbs := int(math.Abs(float64(dividend)))
	divisorAbs := int(math.Abs(float64(divisor)))
	result := div(dividendAbs, divisorAbs)

	if (dividend <= 0 && divisor > 0) || (dividend >= 0 && divisor < 0) {
		return -result
	}
	return result
}

func div(dividend int, divisor int) int {
	result := 0
	for dividend >= divisor {
		multi := 1
		for multi*divisor <= (dividend >> 1) {
			multi <<= 1
		}
		result += multi
		// 相减的结果进入下次循环
		dividend -= multi * divisor
	}
	return result
}
