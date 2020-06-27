package leetcode

func exchange(nums []int) []int {
	head, tail := 0, len(nums)-1
	for head < tail {
		for head < tail && nums[head]&1 != 0 {
			head++
		}
		for head < tail && nums[tail]&1 == 0 {
			tail--
		}
		nums[head], nums[tail] = nums[tail], nums[head]
		head++
		tail--
	}
	return nums
}
