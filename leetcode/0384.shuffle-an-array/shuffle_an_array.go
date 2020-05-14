package leetcode

import "math/rand"

// Solution solution
type Solution struct {
	nums []int
}

// Constructor initialize your solution
func Constructor(nums []int) Solution {
	return Solution{nums}
}

// Reset resets the array to its original configuration and return it.
func (s *Solution) Reset() []int {
	return s.nums
}

// Shuffle returns a random shuffling of the array.
func (s *Solution) Shuffle() []int {
	nums := make([]int, len(s.nums))
	copy(nums, s.nums)
	rand.Shuffle(len(nums), func(i int, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}
