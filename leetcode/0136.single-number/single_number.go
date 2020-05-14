package leetcode

func singleNumber(nums []int) int {
	a := 0
	for _, i := range nums {
		a ^= i
	}
	return a
}
