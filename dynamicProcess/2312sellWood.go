package dynamicprocess

func sellingWood(m int, n int, prices [][]int) int64 {
	pr := make([][]int, m+1)
	for i := range pr {
		pr[i] = make([]int, n+1)
	}
	for _, price := range prices {
		pr[price[0]][price[1]] = price[2]
	}
	dp := make([][]int64, m+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = int64(pr[i][j])
			for k := 1; k < j; k++ {
				dp[i][j] = max64(dp[i][j], dp[i][k]+dp[i][j-k])
			}
			for k := 1; k < i; k++ {
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
