package leetcode

func shuffle(nums []int, n int) []int {
	result := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		result = append(result, nums[i], nums[i+n])
	}
	return result
}
