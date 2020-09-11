package leetcode

var ans [][]int

func subsets(nums []int) [][]int {
	ans = make([][]int, 0)

	for i := 0; i <= len(nums); i++ {
		cur := make([]int, 0, i)
		backtrace(nums, 0, i, cur)
	}

	return ans
}

func backtrace(nums []int, pos, num int, cur []int) {
	if len(cur) == num {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		ans = append(ans, tmp)
		return
	}

	for i := pos; i < len(nums); i++ {
		cur = append(cur, nums[i])
		backtrace(nums, i+1, num, cur)
		cur = cur[:len(cur)-1]
	}
}
