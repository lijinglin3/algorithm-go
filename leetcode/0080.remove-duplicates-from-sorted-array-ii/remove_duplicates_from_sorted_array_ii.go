package leetcode

func removeDuplicates(nums []int) int {
	i, j := 1, 2
	for ; j < len(nums); j++ {
		if nums[j] != nums[i-1] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
