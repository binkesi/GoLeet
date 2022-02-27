package dynamicprocess

// https://leetcode-cn.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	wDict := map[string]struct{}{}
	lenS := len(s)
	for _, v := range wordDict {
		wDict[v] = struct{}{}
	}
	dp := make([]bool, lenS)
	for i := 0; i < lenS; i++ {
		for j := i - 1; j >= -1; j-- {
			if j == -1 || dp[j] {
				tmp := s[j+1 : i+1]
				if _, ok := wDict[tmp]; ok {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[lenS-1]
}
