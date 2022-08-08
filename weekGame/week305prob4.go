package weekgame

// https://leetcode.cn/contest/weekly-contest-305/problems/longest-ideal-subsequence/

func longestIdealString(s string, k int) (ans int) {
	l := len(s)
	dp := make([]int, 26)
	for i := 0; i < l; i++ {
		ind := int(s[i] - 'a')
		dp[ind]++
		for j, v := range dp {
			if j != ind && v != 0 && abs(ind-j) <= k {
				res := v + 1
				if res > dp[ind] {
					dp[ind] = res
				}
			}
		}
	}
	for _, v := range dp {
		if v > ans {
			ans = v
		}
	}
	return
}
