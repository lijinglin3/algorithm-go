package leetcode

func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else if count > 0 {
			nums[i-count] = nums[i]
			nums[i] = 0
		}
	}
}
