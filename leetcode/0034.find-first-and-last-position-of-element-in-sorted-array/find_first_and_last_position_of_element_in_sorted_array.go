package leetcode

func searchRange(nums []int, target int) []int {
	l, r := -1, -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			for j := mid; j < len(nums) && nums[j] <= target; j++ {
				r = j
			}
			for j := mid; j >= 0 && nums[j] == target; j-- {
				l = j
			}
			break
		}
	}
	return []int{l, r}
}
