package twopointer

// https://leetcode.cn/problems/di-string-match/submissions/

func diStringMatch(s string) (ans []int) {
	n := len(s)
	l, h := 0, n
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			ans = append(ans, l)
			l++
		} else {
			ans = append(ans, h)
			h--
		}
	}
	ans = append(ans, h)
	return
}
