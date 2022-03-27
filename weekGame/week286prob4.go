package weekgame

// https://leetcode-cn.com/problems/maximum-value-of-k-coins-from-piles/

func maxValueOfCoins(piles [][]int, k int) (ans int) {
	n := len(piles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	sumN := min(len(piles[0]), k)
	for i := 1; i <= sumN; i++ {
		dp[0][i] = dp[0][i-1] + piles[0][i-1]
	}
	for i := 1; i < n; i++ {
		l := len(piles[i])
		for j := 1; j < l; j++ {
			piles[i][j] += piles[i][j-1]
		}
		sumN = min(l+sumN, k)
		for c := sumN; c > 0; c-- {
			dp[i][c] = dp[i-1][c]
			for w, v := range piles[i][:min(l, c)] {
				dp[i][c] = maxN(dp[i][c], dp[i-1][c-w-1]+v)
			}
		}
	}
	return dp[n-1][k]
}
