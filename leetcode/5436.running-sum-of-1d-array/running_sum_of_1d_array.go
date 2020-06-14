package leetcode

func runningSum(nums []int) []int {
	sum := 0
	for i, v := range nums {
		sum += v
		nums[i] = sum
	}
	return nums
}
