package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	left, right, count := 0, 0, 0
	for ; right < len(nums); right++ {
		if nums[right] == 0 {
			count++
		}
		if count > 1 {
			if nums[left] == 0 {
				count--
			}
			left++
		}

	}
	return right - left
}
