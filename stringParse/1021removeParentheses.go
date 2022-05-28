package stringparse

import "strings"

// https://leetcode.cn/problems/remove-outermost-parentheses/submissions/

func removeOuterParentheses(s string) string {
	n := len(s)
	var b1 strings.Builder
	c1, c2, start := 0, 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			c1++
		}
		if s[i] == ')' {
			c2++
		}
		if c1 == c2 && c1 != 0 {
			b1.WriteString(s[start+1 : i])
			start = i + 1
			c1 = 0
			c2 = 0
		}
	}
	return b1.String()
}
