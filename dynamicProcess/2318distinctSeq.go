package dynamicprocess

// https://leetcode.cn/problems/number-of-distinct-roll-sequences/submissions/

const mod int = 1e9 + 7

func distinctSequences(n int) (ans int) {
	if n == 1 {
		return 6
	}
	if n == 2 {
		return 22
	}
	dp := make([][7][7]int, n)
	dp[0][0][1], dp[0][0][2], dp[0][0][3], dp[0][0][4], dp[0][0][5], dp[0][0][6] = 1, 1, 1, 1, 1, 1
	for j := 1; j <= 6; j++ {
		for i := 1; i <= 6; i++ {
			if j != i && div(j, i) == 1 {
				dp[1][i][j] += dp[0][0][i]
			}
		}
	}
	for k := 2; k < n; k++ {
		for j := 1; j <= 6; j++ {
			for i := 1; i <= 6; i++ {
				for s := 1; s <= 6; s++ {
					if j != i && j != s && div(j, i) == 1 {
						dp[k][i][j] = (dp[k][i][j] + dp[k-1][s][i]) % mod
					}
				}
			}
		}
	}
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			ans = (ans + dp[n-1][i][j]) % mod
		}
	}
	return ans
}

func div(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
