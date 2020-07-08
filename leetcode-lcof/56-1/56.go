package leetcode

func singleNumbers(nums []int) []int {
	var a int
	for i := range nums {
		a ^= nums[i]
	}
	mask := a & (-a)
	res := make([]int, 2)
	for _, v := range nums {
		if (v & mask) == 0 {
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}
	return res
}
