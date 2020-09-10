package leetcode

var result [][]int

func permute(nums []int) [][]int {
	result = make([][]int, 0)
	pathNums := make([]int, 0)
	used := make([]bool, len(nums))
	backtrack(nums, pathNums, used)
	return result
}

func backtrack(nums, pathNums []int, used []bool) {
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(nums))
		copy(tmp, pathNums)
		result = append(result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			pathNums = append(pathNums, nums[i])
			backtrack(nums, pathNums, used)
			pathNums = pathNums[:len(pathNums)-1]
			used[i] = false
		}
	}
}
