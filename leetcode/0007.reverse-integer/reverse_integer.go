package leetcode

import (
	"math"
)

func reverse(x int) int {
	v := 0
	for x != 0 {
		v = v*10 + x%10
		x = x / 10
	}
	if v > math.MaxInt32 || v < math.MinInt32 {
		return 0
	}
	return v
}
