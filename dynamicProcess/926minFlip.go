package dynamicprocess

// https://leetcode.cn/problems/flip-string-to-monotone-increasing/

func minFlipsMonoIncr(s string) int {
	n := len(s)
	dp := make([][2]int, n)
	if s[0] == '0' {
		dp[0][0] = 0
		dp[0][1] = 1
	}
	if s[0] == '1' {
		dp[0][0] = 1
		dp[0][1] = 0
	}
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
		}
		if s[i] == '1' {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		}
	}
	return min(dp[n-1][0], dp[n-1][1])
}
