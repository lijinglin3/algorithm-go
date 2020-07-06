package leetcode

func isStraight(nums []int) bool {
	count, min, max, m := 0, 14, 0, make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else {
			if m[nums[i]] {
				return false
			}
			m[nums[i]] = true
			if max < nums[i] {
				max = nums[i]
			}
			if min > nums[i] {
				min = nums[i]
			}
		}
	}
	if count == 0 {
		return max-min == 4
	}
	return max-min <= 4
}
