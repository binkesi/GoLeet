package dynamicprocess

func longestContnuousSubStr(text1, text2 string) (ans string) {
	m, n := len(text1), len(text2)
	// 这里的dp定义为以text1[m-1]和text2[n-1]结尾的最长连续公共子串的长度
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				l := dp[i-1][j-1] + 1
				if l > len(ans) {
					if j != n {
						ans = text2[j-l : j]
					} else {
						ans = text2[j-l:]
					}
				}
				dp[i][j] = l
			} else {
				dp[i][j] = 0
			}
		}
	}
	return
}
