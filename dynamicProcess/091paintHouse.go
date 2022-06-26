package dynamicprocess

// https://leetcode.cn/problems/JEj789/submissions/

func minCost(costs [][]int) (ans int) {
	n := len(costs)
	dp := make([][3]int, n)
	dp[0][0] = costs[0][0]
	dp[0][1] = costs[0][1]
	dp[0][2] = costs[0][2]
	for i := 1; i < n; i++ {
		for j := 0; j < 3; j++ {
			dp[i][j] = min(dp[i-1][(j+1)%3], dp[i-1][(j+2)%3]) + costs[i][j]
		}
	}
	ans = dp[n-1][0]
	for j := 1; j < 3; j++ {
		if dp[n-1][j] < ans {
			ans = dp[n-1][j]
		}
	}
	return
}
