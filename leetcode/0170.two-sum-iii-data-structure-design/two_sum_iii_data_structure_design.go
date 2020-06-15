package leetcode

import "math"

// TwoSum ...
type TwoSum struct {
	val map[int]int
	max int64
	min int64
}

// Constructor initialize your data structure here.
func Constructor() TwoSum {
	return TwoSum{val: make(map[int]int), min: math.MaxInt64, max: math.MinInt64}

}

// Add the number to an internal data structure.
func (s *TwoSum) Add(number int) {
	if s.max < int64(number) {
		s.max = int64(number)
	}
	if s.min > int64(number) {
		s.min = int64(number)
	}

	s.val[number]++
}

// Find if there exists any pair of numbers which sum is equal to the value.
func (s *TwoSum) Find(value int) bool {

	if int64(value) < s.min+s.min || int64(value) > s.max+s.max {
		return false
	}
	for k, v := range s.val {
		if value-k == k {
			if v > 1 {
				return true
			}

		} else {
			if s.val[value-k] > 0 {
				return true
			}
		}
	}
	return false
}
