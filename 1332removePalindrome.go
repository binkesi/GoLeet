package goleet

// https://leetcode-cn.com/problems/remove-palindromic-subsequences/

func removePalindromeSub(s string) int {
	lenS := len(s)
	for i, j := 0, lenS-1; i < j; {
		if s[i] != s[j] {
			return 2
		}
		i += 1
		j -= 1
	}
	return 1
}
