package greedy

import "math"

// https://leetcode.cn/problems/longest-binary-subsequence-less-than-or-equal-to-k/submissions/

func longestSubsequence(s string, k int) (ans int) {
	l := len(s)
	res := 0
	for i := l - 1; i >= 0; i-- {
		if s[i] == '0' {
			ans++
		} else {
			if ans >= 63 { // 注意大于等于64位的1一定会超过k
				continue
			}
			num := math.Pow(2.0, float64(ans))
			if res+int(num) <= k {
				res += int(num)
				ans++
			}
		}
	}
	return
}
