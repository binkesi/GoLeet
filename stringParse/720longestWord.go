package stringparse

// https://leetcode-cn.com/problems/longest-word-in-dictionary/

func longestWord(words []string) (ans string) {
	maxLen := 0
	wMap := make(map[string]struct{})
	for _, w := range words {
		wMap[w] = struct{}{}
	}
next:
	for _, w := range words {
		n := len(w)
		for i := 1; i < n; i++ {
			if _, ok := wMap[w[:i]]; !ok {
				continue next
			}
		}
		if n > maxLen {
			maxLen = n
			ans = w
		}
		if n == maxLen {
			for i := 0; i < n; i++ {
				if w[i] < ans[i] {
					ans = w
					break
				} else if w[i] > ans[i] {
					break
				}
			}
		}
	}
	return
}
