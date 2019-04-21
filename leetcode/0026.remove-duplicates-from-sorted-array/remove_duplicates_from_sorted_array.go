package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	for i := 0; ; i++ {
		if i == len(nums)-1 {
			return len(nums)
		}
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
}
