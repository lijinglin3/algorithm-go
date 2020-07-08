package leetcode

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var counts = make([]int, 32)
	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		for j := 0; j < 32; j++ {
			counts[j] = counts[j] + tmp&1
			tmp = tmp >> 1
		}
	}
	var res = 0
	var m = 3
	for i := 0; i < 32; i++ {
		res = res << 1
		res = res | counts[31-i]%m
	}
	return res
}
