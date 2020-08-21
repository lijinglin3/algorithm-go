package leetcode

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}
	var result []int
	for k, v := range nums {
		if v > 0 {
			result = append(result, k+1)
		}
	}
	return result
}
