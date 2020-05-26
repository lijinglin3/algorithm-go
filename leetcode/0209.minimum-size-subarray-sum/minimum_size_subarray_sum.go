package leetcode

import "math"

func minSubArrayLen(s int, nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	count := math.MaxInt64
	sum := nums[0]
	p1, p2 := 0, 0
	for {
		if sum < s {
			if p2 < len(nums)-1 {
				p2++
				sum += nums[p2]
			} else {
				break
			}
		} else {
			if p1 < len(nums)-1 {
				if count > p2-p1+1 {
					count = p2 - p1 + 1
				}
				sum -= nums[p1]
				p1++
			} else {
				break
			}
		}
	}
	if count == math.MaxInt64 {
		return 0
	}
	return count
}
