package stringparse

// https://leetcode.cn/problems/find-and-replace-pattern/submissions/

func findAndReplacePattern(words []string, pattern string) (ans []string) {
	n := len(pattern)
	for _, word := range words {
		ref := make(map[byte]byte)
		rev := make(map[byte]byte)
		for i, b := range pattern {
			if _, ok := ref[byte(b)]; !ok {
				if _, ok := rev[word[i]]; !ok {
					ref[byte(b)] = word[i]
					rev[word[i]] = byte(b)
				} else {
					break
				}
			} else {
				if ref[byte(b)] != word[i] || rev[word[i]] != byte(b) {
					break
				}
			}
			if i == n-1 {
				ans = append(ans, word)
			}
		}
	}
	return
}
