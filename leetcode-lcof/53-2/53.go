package leetcode

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if mid == nums[mid] {
			left = mid + 1
		}
		if mid < nums[mid] {
			right = mid - 1
		}
	}
	return left
}
