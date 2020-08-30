package leetcode

import "math"

func increasingTriplet(nums []int) bool {
	m1, m2 := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if m1 >= v {
			m1 = v
		} else if m2 >= v {
			m2 = v
		} else {
			return true
		}
	}
	return false
}
