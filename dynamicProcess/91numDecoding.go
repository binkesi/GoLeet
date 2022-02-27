package dynamicprocess

// https://leetcode-cn.com/problems/decode-ways/

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	lenS := len(s)
	dp := make([]int, lenS)
	dp[0] = 1
	for i := 1; i < lenS; i++ {
		if s[i] == '0' {
			if s[i-1]-'0' > 0 && s[i-1]-'0' <= 2 {
				if i > 1 {
					dp[i] = dp[i-2]
				}
				if i == 1 {
					dp[i] = dp[i-1]
				}
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			if i > 1 {
				dp[i] = dp[i-1] + dp[i-2]
			}
			if i == 1 {
				dp[i] = dp[i-1] + 1
			}
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[lenS-1]
}
