package leetcode

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, 0, len(nums))
	max := 1
	for i, v1 := range nums {
		dp = append(dp, 1)
		for j, v2 := range nums[:i] {
			if v1 > v2 {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					if max < dp[i] {
						max = dp[i]
					}
				}
			}
		}
	}
	return max
}
