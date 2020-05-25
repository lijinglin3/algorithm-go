package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	result, count := 0, 0
	for _, v := range nums {
		if v == 1 {
			count++
		} else {
			if count > result {
				result = count
			}
			count = 0
		}
	}
	if count > result {
		result = count
	}
	return result
}
