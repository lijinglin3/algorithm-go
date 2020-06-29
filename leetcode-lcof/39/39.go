package leetcode

func majorityElement(nums []int) int {
	x, votes := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if votes == 0 {
			x = nums[i]
		}

		if nums[i] == x {
			votes++
		} else {
			votes--
		}
	}
	return x
}
