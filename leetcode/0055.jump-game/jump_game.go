package leetcode

func canJump(nums []int) bool {
	k := 0
	for i, v := range nums {
		if i > k {
			return false
		}
		if k < i+v {
			k = i + v
		}
	}
	return true
}
