package dynamicprocess

// https://leetcode-cn.com/problems/push-dominoes/

func pushDominoes(dominoes string) string {
	lenD := len(dominoes)
	if lenD == 1 {
		return dominoes
	}
	dp := make([]byte, lenD)
	for i := 0; i < lenD; i++ {
		dp[i] = dominoes[i]
	}
	for {
		change := false
		tmp := make([]byte, lenD)
		copy(tmp, dp)
		for j := 0; j < lenD; j++ {
			if j == 0 {
				if tmp[j] == '.' && dp[j+1] == 'L' {
					dp[j] = 'L'
					change = true
				}
				continue
			}
			if j == lenD-1 {
				if tmp[j] == '.' && dp[j-1] == 'R' {
					dp[j] = 'R'
					change = true
				}
				continue
			}
			if tmp[j] == '.' {
				if (tmp[j-1] == '.' || tmp[j-1] == 'L') && tmp[j+1] == 'L' {
					dp[j] = 'L'
					change = true
				} else if tmp[j-1] == 'R' && (tmp[j+1] == '.' || tmp[j+1] == 'R') {
					dp[j] = 'R'
					change = true
				}
			}
		}
		if !change {
			break
		}
	}
	return string(dp)
}
