package weekgame

// https://leetcode-cn.com/contest/weekly-contest-258/problems/maximum-product-of-the-length-of-two-palindromic-subsequences/

func maxProduct(s string) int {
	lenS := len(s)
	ans := 1
	for i := 1; i < (1<<lenS)-1; i++ {
		var a, b string
		for j := 0; j < lenS; j++ {
			if (i>>j)&1 == 1 {
				a += string(s[j])
			} else {
				b += string(s[j])
			}
		}
		ans = max2(ans, maxLength(a)*maxLength(b))
	}
	return ans
}

func maxLength(s string) int {
	lenS := len(s)
	dp := make([][]int, lenS)
	for i := range dp {
		dp[i] = make([]int, lenS)
	}
	for i := 0; i < lenS; i++ {
		dp[i][i] = 1
	}
	for i := 0; i < lenS-1; i++ {
		for j := i + 1; j < lenS; j++ {
			if s[i] == s[j] {
				dp[i][j] = max2(dp[i+1][j-1]+2, dp[i][j])
			} else {
				dp[i][j] = max2(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][lenS-1]
}

func max2(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
