package leetcode

import "math"

func reverse(x int) int {
	v := 0
	for x != 0 {
		v = v*10 + x%10
		x = x / 10
	}
	// 判断溢出
	if v-math.MaxInt32 > 0 || v-math.MinInt32 < 0 {
		v = 0
	}
	return v
}
