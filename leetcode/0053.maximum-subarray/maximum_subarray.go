package leetcode

func maxSubArray(nums []int) int {
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > sum {
			sum = nums[i]
		}
	}
	return sum
}

// TODO 尝试分治法
