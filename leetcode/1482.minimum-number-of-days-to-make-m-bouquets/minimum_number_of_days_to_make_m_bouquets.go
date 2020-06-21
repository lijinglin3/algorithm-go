package leetcode

import (
	"math"
)

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	max := 0
	res := -1
	min := int(math.Pow10(9))
	for i := 0; i < len(bloomDay); i++ {
		if bloomDay[i] > max {
			max = bloomDay[i]
		}
		if bloomDay[i] < min {
			min = bloomDay[i]
		}
	}
	l := min
	r := max
	for l <= r {
		mid := (l + r) / 2
		if check(bloomDay, m, k, mid) {
			r = mid - 1
			res = mid
		} else {
			l = mid + 1
		}

	}
	return res
}

func check(b []int, m int, k int, t int) bool {

	cur := 0
	for i := 0; i < len(b); i++ {
		if b[i] <= t {
			cur++
		} else {
			m -= cur / k
			cur = 0
		}
	}
	m -= cur / k
	return m <= 0
}
