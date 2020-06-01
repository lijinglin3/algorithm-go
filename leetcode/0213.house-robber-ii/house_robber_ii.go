package leetcode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp1 := dp(nums[:len(nums)-1])
	dp2 := dp(nums[1:])

	if dp1 > dp2 {
		return dp1
	}
	return dp2
}

func dp(nums []int) int {
	tmp := []int{0, 0}
	for i := 2; i < len(nums)+2; i++ {
		if tmp[i-1] > tmp[i-2]+nums[i-2] {
			tmp = append(tmp, tmp[i-1])
		} else {
			tmp = append(tmp, tmp[i-2]+nums[i-2])
		}
	}
	return tmp[len(tmp)-1]
}
