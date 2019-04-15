package leetcode

import (
	"math"
	"testing"
)

// https://leetcode-cn.com/problems/reverse-integer/

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

func TestReverse(t *testing.T) {
	if reverse(123) != 321 {
		t.Fail()
	}
	if reverse(-120) != -21 {
		t.Fail()
	}
	if reverse(0) != 0 {
		t.Fail()
	}
	if reverse(1534236469) != 0 {
		t.Fail()
	}
}
