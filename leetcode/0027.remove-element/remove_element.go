package leetcode

func removeElement(nums []int, val int) int {
	for i := 0; ; i++ {
		if i == len(nums) {
			return len(nums)
		}
		if nums[i] == val {
			nums[i] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			i--
		}
	}
}
