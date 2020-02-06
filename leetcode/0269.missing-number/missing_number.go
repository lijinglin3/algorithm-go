package leetcode

func missingNumber(nums []int) int {
	result := len(nums)
	for i, v := range nums {
		result ^= i ^ v
	}
	return result
}
