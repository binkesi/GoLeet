package weekgame

// https://leetcode-cn.com/contest/weekly-contest-219/problems/stone-game-vii/

func stoneGameVII2(stones []int) int {
	lenS := len(stones)
	sums := make([]int, lenS+1)
	for i := 1; i < lenS+1; i++ {
		sums[i] = sums[i-1] + stones[i-1]
	}
	dp := make([][]int, lenS)
	for i := 0; i < lenS; i++ {
		dp[i] = make([]int, lenS)
	}
	for i := lenS - 2; i >= 0; i-- {
		for j := i + 1; j < lenS; j++ {
			dp[i][j] = maxTwo(sums[j+1]-sums[i+1]-dp[i+1][j], sums[j]-sums[i]-dp[i][j-1])
		}
	}
	return dp[0][lenS-1]
}

func maxTwo(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
