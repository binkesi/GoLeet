package arraymanipulate

import "unicode"

// https://leetcode-cn.com/problems/reverse-only-letters/

func reverseOnlyLetters(s string) string {
	lenS := len(s)
	sbyte := []byte(s)
	for i, j := 0, lenS-1; i < j; {
		l, r := sbyte[i], sbyte[j]
		if unicode.IsLetter(rune(l)) && unicode.IsLetter(rune(r)) {
			sbyte[i], sbyte[j] = sbyte[j], sbyte[i]
			i += 1
			j -= 1
		}
		if !unicode.IsLetter(rune(l)) {
			i += 1
		}
		if !unicode.IsLetter(rune(r)) {
			j -= 1
		}
	}
	return string(sbyte)
}
