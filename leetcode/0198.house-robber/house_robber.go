package leetcode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := []int{0, nums[0]}
	for i := 2; i <= len(nums); i++ {
		if dp[i-1] > dp[i-2]+nums[i-1] {
			dp = append(dp, dp[i-1])
		} else {
			dp = append(dp, dp[i-2]+nums[i-1])
		}
	}
	return dp[len(nums)]
}
