package leetcode

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for nums[left] >= nums[right] && left < right {
		mid := (left + right) / 2
		if nums[left] < nums[mid] {
			left = mid + 1
		} else if nums[left] > nums[mid] {
			right = mid
		} else {
			left++
		}
	}
	return nums[left]
}
