package arraymanipulate

// https://leetcode-cn.com/problems/reverse-only-letters/

func reverseOnlyLetters(s string) string {
	lenS := len(s)
	sbyte := []byte(s)
	for i, j := 0, lenS-1; i < j; {
		l, r := sbyte[i], sbyte[j]
		if isLetter(l) && isLetter(r) {
			sbyte[i], sbyte[j] = sbyte[j], sbyte[i]
			i += 1
			j -= 1
		}
		if !isLetter(l) {
			i += 1
		}
		if !isLetter(r) {
			j -= 1
		}
	}
	return string(sbyte)
}

func isLetter(a byte) bool {
	return (a-'a' >= 0 && a-'a' < 26) || (a-'A' >= 0 && a-'A' < 26)
}
