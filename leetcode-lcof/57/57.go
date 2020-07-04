package leetcode

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		tmp := nums[left] + nums[right]
		switch {
		case tmp > target:
			right--
		case tmp < target:
			left++
		default:
			return []int{nums[left], nums[right]}
		}
	}
	return nil
}
