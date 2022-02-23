package dynamicprocess

// https://leetcode-cn.com/problems/word-break-ii/

func wordBreak2(s string, wordDict []string) (ans []string) {
	lenS := len(s)
	dp := make([][]string, lenS)
	wDict := make(map[string]struct{})
	for _, v := range wordDict {
		wDict[v] = struct{}{}
	}
	for i := 0; i < lenS; i++ {
		dp[i] = []string{}
		for j := i - 1; j >= -1; j-- {
			str := s[j+1 : i+1]
			if j == -1 {
				if _, ok := wDict[str]; ok {
					dp[i] = append(dp[i], str)
				}
			} else if len(dp[j]) != 0 {
				if _, ok := wDict[str]; ok {
					for _, v := range dp[j] {
						dp[i] = append(dp[i], v+" "+str)
					}
				}
			}
		}
	}
	ans = dp[lenS-1]
	return
}
