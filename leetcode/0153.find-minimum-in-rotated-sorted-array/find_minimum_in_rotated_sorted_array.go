package leetcode

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end-start)>>1
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = end - 1
		}
	}
	return nums[start]
}
