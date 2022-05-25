package dynamicprocess

// https://leetcode.cn/problems/unique-substrings-in-wraparound-string/submissions/

func findSubstringInWraproundString(p string) (ans int) {
	dp := [26]int{}
	k := 0
	for i, v := range p {
		if i > 0 && (byte(v)-p[i-1]+26)%26 == 1 {
			k++
		} else {
			k = 1
		}
		dp[v-'a'] = max(dp[v-'a'], k)
	}
	for _, v := range dp {
		ans += v
	}
	return
}
