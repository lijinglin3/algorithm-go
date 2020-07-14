package leetcode

// NumArray ...
type NumArray struct {
	data []int
}

// Constructor ...
func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	sums[0] = 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sums[i+1] = sum
	}
	return NumArray{data: sums}
}

// SumRange ...
func (a *NumArray) SumRange(i int, j int) int {
	return a.data[j+1] - a.data[i]
}
