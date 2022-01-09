package goleet

func isMatchReg(s string, p string) bool {
	lens, lenp := len(s), len(p)
	var dp [][]bool = make([][]bool, lens+1)
	for a := range dp {
		dp[a] = make([]bool, lenp+1)
	}
	dp[0][0] = true
	for i := 1; i <= lens; i++ {
		dp[i][0] = false
	}
	for j := 1; j <= lenp; j++ {
		if p[j-1] != '*' {
			dp[0][j] = false
		} else if dp[0][j-1] == true {
			dp[0][j] = true
		}
	}
	for l := 1; l <= lens; l++ {
		for c := 1; c <= lenp; c++ {
			if p[c-1] == '?' || p[c-1] == s[l-1] {
				dp[l][c] = dp[l-1][c-1]
			} else if p[c-1] == '*' {
				if !dp[l][c-1] && !dp[l-1][c] {
					dp[l][c] = false
				} else {
					dp[l][c] = true
				}
			} else {
				dp[l][c] = false
			}
		}
	}
	return dp[lens][lenp]
}
