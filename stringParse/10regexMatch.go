package stringparse

// https://leetcode-cn.com/problems/regular-expression-matching/

func isMatchOk(s string, p string) bool {
	lens, lenp := len(s), len(p)
	var dp [][]bool = make([][]bool, lens+1)
	for i := range dp {
		dp[i] = make([]bool, lenp+1)
	}
	dp[0][0] = true
	for i := 1; i <= lens; i++ {
		dp[i][0] = false
	}
	for j := 1; j <= lenp; j++ {
		if p[j-1] != '*' {
			dp[0][j] = false
		} else {
			dp[0][j] = dp[0][j-2]
		}
	}
	for l := 1; l <= lens; l++ {
		for c := 1; c <= lenp; c++ {
			if p[c-1] == '.' || p[c-1] == s[l-1] {
				dp[l][c] = dp[l-1][c-1]
			} else if p[c-1] == '*' {
				if p[c-2] == '.' {
					for k := 0; k <= l; k++ {
						if dp[k][c-2] == true {
							dp[l][c] = true
							break
						}
						if k == l {
							dp[l][c] = false
						}
					}
				} else {
					dp[l][c] = dp[l][c-2] || (dp[l-1][c-2] && p[c-2] == s[l-1]) || (dp[l-1][c] && p[c-2] == s[l-1])
				}
			} else {
				dp[l][c] = false
			}
		}
	}
	return dp[lens][lenp]
}
