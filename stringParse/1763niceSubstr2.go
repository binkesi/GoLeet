package stringparse

import "unicode"

// https://leetcode-cn.com/problems/longest-nice-substring/

func longestNiceSubstring2(s string) string {
	var ans string
	lenS := len(s)
	for i := range s {
		lower, upper := 0, 0
		for j := i; j < lenS; j++ {
			if unicode.IsLower(rune(s[j])) {
				lower |= (1 << (s[j] - 'a'))
			} else {
				upper |= (1 << (s[j] - 'A'))
			}
			if lower == upper && (j-i+1) > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}
