package stringparse

import "unicode"

// https://leetcode-cn.com/problems/longest-nice-substring/

func longestNiceSubstring3(s string) (ans string) {
	if s == "" {
		return
	}
	lower, upper := 0, 0
	for _, ch := range s {
		if unicode.IsLower(rune(ch)) {
			lower |= (1 << (ch - 'a'))
		} else {
			upper |= (1 << (ch - 'A'))
		}
	}
	if lower == upper {
		return s
	}
	valid := lower & upper
	for i := 0; i < len(s); i++ {
		start := i
		for i < len(s) && valid>>(unicode.ToLower(rune(s[i]))-'a')&1 == 1 {
			i += 1
		}
		if t := longestNiceSubstring3(s[start:i]); len(t) > len(ans) {
			ans = t
		}
	}
	return
}
