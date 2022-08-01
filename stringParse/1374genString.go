package stringparse

import "strings"

// https://leetcode.cn/problems/generate-a-string-with-characters-that-have-odd-counts/

func generateTheString(n int) string {
	res := strings.Builder{}
	if n%2 == 1 {
		for i := 0; i < n; i++ {
			res.WriteByte('a')
		}
	} else {
		for i := 0; i < n-1; i++ {
			res.WriteByte('a')
		}
		res.WriteByte('b')
	}
	return res.String()
}
