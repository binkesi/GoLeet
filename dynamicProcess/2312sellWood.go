package dynamicprocess

// https://leetcode.cn/problems/selling-pieces-of-wood/

func sellingWood(m int, n int, prices [][]int) int64 {
	dp := make([][]int64, m+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}
	for _, price := range prices {
		dp[price[0]][price[1]] = int64(price[2])
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= j/2; k++ {
				dp[i][j] = max64(dp[i][j], dp[i][k]+dp[i][j-k])
			}
			for k := 1; k <= i/2; k++ {
				dp[i][j] = max64(dp[i][j], dp[k][j]+dp[i-k][j])
			}
		}
	}
	return dp[m][n]
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
