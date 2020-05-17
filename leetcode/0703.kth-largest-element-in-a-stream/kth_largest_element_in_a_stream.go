package leetcode

import "sort"

// KthLargest KthLargest
type KthLargest struct {
	nums []int
	k    int
}

// Constructor initial KthLargest
func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	if len(nums)-k > 0 {
		nums = nums[len(nums)-k:]
	}
	return KthLargest{
		nums: nums,
		k:    k,
	}
}

// Add return Kth largest number
func (l *KthLargest) Add(val int) int {
	if len(l.nums) == 0 {
		l.nums = append(l.nums, val)
		return val
	}

	if l.nums[len(l.nums)-1] < val {
		l.nums = append(l.nums, val)
	} else {
		for i := 0; i < len(l.nums); i++ {
			if l.nums[i] >= val {
				l.nums = append(l.nums, val)
				copy(l.nums[i+1:], l.nums[i:])
				l.nums[i] = val
				break
			}
		}
	}

	if len(l.nums) > l.k {
		l.nums = l.nums[1:]
	}

	return l.nums[0]
}
