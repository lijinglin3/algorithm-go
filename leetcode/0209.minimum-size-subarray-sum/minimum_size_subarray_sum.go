package leetcode

import (
	"math"
)

func minSubArrayLen(s int, nums []int) int {
	length := len(nums)
	if length <= 0 {
		return 0
	}

	count, start, end, sum := math.MaxInt64, 0, 0, 0
	for end < length {
		sum += nums[end]
		for sum >= s {
			if count > (end - start + 1) {
				count = end - start + 1
			}
			sum -= nums[start]
			start++
		}
		end++
	}
	if count == math.MaxInt64 {
		return 0
	}
	return count
}
