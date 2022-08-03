package stringparse

import (
	"sort"
)

// https://leetcode.cn/problems/orderly-queue/

func orderlyQueue(s string, k int) string {
	l, ans := len(s), s
	if l == 1 {
		return s
	}
	b := []byte(s)
	if k != 1 {
		sort.Slice(b, func(i, j int) bool {
			return b[i]-'a' < b[j]-'a'
		})
		return string(b)
	}
	for i := 0; i < l; i++ {
		s = s[1:] + s[:1]
		if s < ans {
			ans = s
		}
	}
	return ans
}
