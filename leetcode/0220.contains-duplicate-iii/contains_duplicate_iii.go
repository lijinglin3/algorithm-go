package leetcode

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	length := len(nums)
	for i := 1; i <= k; i++ {
		for j := 0; j < length-i; j++ {
			tmp := nums[j] - nums[j+i]
			if tmp < 0 {
				tmp = -tmp
			}
			if tmp <= t {
				return true
			}
		}
	}
	return false
}

// TODO 尝试其他方法
