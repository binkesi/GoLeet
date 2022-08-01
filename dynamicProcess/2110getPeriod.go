package dynamicprocess

// https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock/submissions/

func getDescentPeriods(prices []int) int64 {
	ans, l := 1, len(prices)
	dp := make([]int, l)
	dp[0] = 1
	for i := 1; i < l; i++ {
		if prices[i]+1 == prices[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		ans += dp[i]
	}
	return int64(ans)
}
