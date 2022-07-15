package dynamicprocess

// https://leetcode.cn/problems/count-number-of-texts/submissions/

func countTexts(pressedKeys string) (ans int) {
	l := len(pressedKeys)
	if l == 1 {
		return 1
	}
	dp := make([]int, l)
	dp[0] = 1
	for i := 1; i < l; i++ {
		dp[i] = dp[i-1]
		if pressedKeys[i] != pressedKeys[i-1] {
			continue
		}
		if pressedKeys[i] == '9' || pressedKeys[i] == '7' {
			for j := i - 1; j >= 0 && j >= i-3; j-- {
				if pressedKeys[j] != pressedKeys[i] {
					break
				} else {
					if j == 0 {
						dp[i] = (dp[i] + 1) % mod
					} else {
						dp[i] = (dp[i] + dp[j-1]) % mod
					}
				}
			}
		} else {
			for j := i - 1; j >= 0 && j >= i-2; j-- {
				if pressedKeys[j] != pressedKeys[i] {
					break
				} else {
					if j == 0 {
						dp[i] = (dp[i] + 1) % mod
					} else {
						dp[i] = (dp[i] + dp[j-1]) % mod
					}
				}
			}
		}
	}
	ans = dp[l-1]
	return
}
