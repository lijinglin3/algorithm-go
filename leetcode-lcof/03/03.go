package leetcode

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if i == n {
			continue
		}
		if nums[n] == n {
			return n
		}
		nums[n], nums[i] = n, nums[n]
	}
	return 0
}
