package leetcode

import "math"

func minMoves(nums []int) int {
	min, sum := math.MaxInt64, 0
	for i := range nums {
		if min > nums[i] {
			min = nums[i]
		}
	}
	for i := range nums {
		sum += nums[i] - min
	}
	return sum
}
