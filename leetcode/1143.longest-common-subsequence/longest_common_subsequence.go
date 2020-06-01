package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	dp := make([][]int, 0, len1+1)
	for i := 0; i < len1+1; i++ {
		dp = append(dp, make([]int, len2+1))
	}

	for i, v1 := range text1 {
		for j, v2 := range text2 {
			if v1 == v2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = dp[i+1][j]
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				}
			}
		}
	}
	return dp[len1][len2]
}
